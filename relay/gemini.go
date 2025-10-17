package relay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/requester"
	"one-api/model"
	providersBase "one-api/providers/base"
	"one-api/providers/gemini"
	"one-api/safty"
	"one-api/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var AllowGeminiChannelType = []int{config.ChannelTypeGemini, config.ChannelTypeVertexAI, config.ChannelTypeOpenAI}

type relayGeminiOnly struct {
	relayBase
	geminiRequest *gemini.GeminiChatRequest
}

func NewRelayGeminiOnly(c *gin.Context) *relayGeminiOnly {
	c.Set("allow_channel_type", AllowGeminiChannelType)
	relay := &relayGeminiOnly{
		relayBase: relayBase{
			allowHeartbeat: true,
			c:              c,
		},
	}

	return relay
}

func (r *relayGeminiOnly) setRequest() error {
	modelAction := r.c.Param("model")

	if modelAction == "" {
		return errors.New("model is required")
	}

	modelList := strings.Split(modelAction, ":")
	if len(modelList) != 2 {
		return errors.New("model error")
	}

	isStream := false
	if modelList[1] == "streamGenerateContent" {
		isStream = true
	}

	r.geminiRequest = &gemini.GeminiChatRequest{}
	if err := common.UnmarshalBodyReusable(r.c, r.geminiRequest); err != nil {
		return err
	}
	r.geminiRequest.Model = modelList[0]
	r.geminiRequest.Stream = isStream
	r.setOriginalModel(r.geminiRequest.Model)

	// 品牌映射校验：仅允许归属Gemini的模型走官方端点
	price := model.PricingInstance.GetPrice(r.geminiRequest.Model)
	if price.ChannelType != config.ChannelTypeGemini {
		return errors.New("模型不属于Gemini品牌，不支持该官方端点")
	}

	return nil
}

func (r *relayGeminiOnly) getRequest() interface{} {
	return r.geminiRequest
}

func (r *relayGeminiOnly) IsStream() bool {
	return r.geminiRequest.Stream
}

func (r *relayGeminiOnly) getPromptTokens() (int, error) {
	channel := r.provider.GetChannel()
	return CountGeminiTokenMessages(r.geminiRequest, channel.PreCost)
}

func (r *relayGeminiOnly) send() (err *types.OpenAIErrorWithStatusCode, done bool) {
	// 优先品牌原生Provider
	chatProvider, ok := r.provider.(gemini.GeminiChatInterface)
	// 不支持则尝试OpenAI Chat回退
	if !ok {
		fallbackChat, ok2 := r.provider.(providersBase.ChatInterface)
		if !ok2 {
			err = common.StringErrorWrapperLocal("channel not implemented", "channel_error", http.StatusServiceUnavailable)
			done = true
			return
		}

		// 转换为OpenAI Chat请求
		openaiReq, convErr := gemini.GeminiToOpenAIChatRequest(r.geminiRequest)
		if convErr != nil {
			err = convErr
			done = true
			return
		}
		openaiReq.Model = r.modelName

		// 安全审查（按OpenAI消息内容）
		if config.EnableSafe {
			for _, m := range openaiReq.Messages {
				if m.Content != nil {
					CheckResult, _ := safty.CheckContent(m.Content)
					if !CheckResult.IsSafe {
						err = common.StringErrorWrapperLocal(CheckResult.Reason, CheckResult.Code, http.StatusBadRequest)
						done = true
						return
					}
				}
			}
		}

		if openaiReq.Stream {
			stream, e := fallbackChat.CreateChatCompletionStream(openaiReq)
			if e != nil {
				err = e
				return
			}
			if r.heartbeat != nil {
				r.heartbeat.Stop()
			}
			first := forwardOpenAIStreamAsGemini(r.c, stream, r.modelName)
			r.SetFirstResponseTime(first)
			return nil, false
		}

		resp, e := fallbackChat.CreateChatCompletion(openaiReq)
		if e != nil {
			err = e
			return
		}
		if r.heartbeat != nil {
			r.heartbeat.Stop()
		}
		gResp := gemini.OpenAIToGeminiChatResponse(resp, r.modelName, r.provider.GetUsage())
		err = responseJsonClient(r.c, gResp)
		if err != nil {
			done = true
		}
		return
	}

	// 内容审查（Gemini原生）
	if config.EnableSafe {
		for _, message := range r.geminiRequest.Contents {
			if message.Parts != nil {
				CheckResult, _ := safty.CheckContent(message.Parts)
				if !CheckResult.IsSafe {
					err = common.StringErrorWrapperLocal(CheckResult.Reason, CheckResult.Code, http.StatusBadRequest)
					done = true
					return
				}
			}
		}
	}

	r.geminiRequest.Model = r.modelName

	if r.geminiRequest.Stream {
		var response requester.StreamReaderInterface[string]
		response, err = chatProvider.CreateGeminiChatStream(r.geminiRequest)
		if err != nil {
			return
		}

		if r.heartbeat != nil {
			r.heartbeat.Stop()
		}

		doneStr := func() string { return "" }
		firstResponseTime := responseGeneralStreamClient(r.c, response, doneStr)
		r.SetFirstResponseTime(firstResponseTime)
	} else {
		var response *gemini.GeminiChatResponse
		response, err = chatProvider.CreateGeminiChat(r.geminiRequest)
		if err != nil {
			return
		}
		if r.heartbeat != nil { r.heartbeat.Stop() }
		err = responseJsonClient(r.c, response)
	}

	if err != nil { done = true }
	return
}

// 将OpenAI流式数据转发为Gemini官方SSE片段
func forwardOpenAIStreamAsGemini(c *gin.Context, stream requester.StreamReaderInterface[string], modelName string) time.Time {
	requester.SetEventStreamHeaders(c)
	dataChan, errChan := stream.Recv()
	done := make(chan struct{})
	defer stream.Close()

	var first time.Time
	var isFirst bool

	go func() {
		defer close(done)
		for {
			select {
			case data, ok := <-dataChan:
				if !ok { return }
				line := strings.TrimSpace(data)
				if strings.HasPrefix(line, "data: ") { line = line[6:] }
				if line == "[DONE]" || line == "{\"id\":\"\",\"object\":\"done\"}" { return }
				var oai types.ChatCompletionStreamResponse
				if err := json.Unmarshal([]byte(line), &oai); err != nil {
					// 非预期片段，透传原始行
					fmtStr := "data: %s\n\n"
					c.Writer.Write([]byte(fmt.Sprintf(fmtStr, line)))
					c.Writer.Flush()
					continue
				}
				// 组装最小Gemini片段（仅文本增量）
				var chunks []gemini.GeminiChatCandidate
				for _, ch := range oai.Choices {
					if ch.Delta.Content == "" { continue }
					chunks = append(chunks, gemini.GeminiChatCandidate{
						Index: int64(ch.Index),
						Content: gemini.GeminiChatContent{Role: "model", Parts: []gemini.GeminiPart{{Text: ch.Delta.Content}}},
					})
				}
				if len(chunks) == 0 { continue }
				g := gemini.GeminiChatResponse{Model: modelName, Candidates: chunks}
				b, _ := json.Marshal(g)
				sse := "data: " + string(b) + "\n\n"
				if !isFirst { first = time.Now(); isFirst = true }
				c.Writer.Write([]byte(sse))
				c.Writer.Flush()
			case err := <-errChan:
				// 出错仍尝试结束
				_ = err
				return
			}
		}
	}()
	<-done
	return first
}

func (r *relayGeminiOnly) GetError(err *types.OpenAIErrorWithStatusCode) (int, any) {
	newErr := FilterOpenAIErr(r.c, err)

	geminiErr := gemini.OpenaiErrToGeminiErr(&newErr)

	return newErr.StatusCode, geminiErr.GeminiErrorResponse
}

func (r *relayGeminiOnly) HandleJsonError(err *types.OpenAIErrorWithStatusCode) {
	statusCode, response := r.GetError(err)
	r.c.JSON(statusCode, response)
}

func (r *relayGeminiOnly) HandleStreamError(err *types.OpenAIErrorWithStatusCode) {
	_, response := r.GetError(err)

	str, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		return
	}
	r.c.Writer.Write([]byte("data: " + string(str) + "\n\n"))
	r.c.Writer.Flush()
}

func CountGeminiTokenMessages(request *gemini.GeminiChatRequest, preCostType int) (int, error) {
	if preCostType == config.PreContNotAll {
		return 0, nil
	}

	tokenEncoder := common.GetTokenEncoder(request.Model)

	tokenNum := 0
	tokensPerMessage := 4
	var textMsg strings.Builder

	for _, message := range request.Contents {
		tokenNum += tokensPerMessage
		for _, part := range message.Parts {
			if part.Text != "" {
				textMsg.WriteString(part.Text)
			}

			if part.InlineData != nil {
				// 其他类型的，暂时按200个token计算
				tokenNum += 200
			}
		}
	}

	if textMsg.Len() > 0 {
		tokenNum += common.GetTokenNum(tokenEncoder, textMsg.String())
	}
	return tokenNum, nil
}
