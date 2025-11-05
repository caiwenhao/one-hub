package relay

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "one-api/common"
    "one-api/common/config"
    "one-api/common/requester"
    "one-api/relay/relay_util"
    "one-api/model"
    providersBase "one-api/providers/base"
    "one-api/providers/gemini"
    "one-api/safty"
    "one-api/types"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "strconv"
)

var AllowGeminiChannelType = []int{config.ChannelTypeGemini, config.ChannelTypeVertexAI, config.ChannelTypeOpenAI}

type relayGeminiOnly struct {
    relayBase
    geminiRequest *gemini.GeminiChatRequest
    action        string
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

    action := strings.TrimSpace(modelList[1])
    isStream := action == "streamGenerateContent"

    r.geminiRequest = &gemini.GeminiChatRequest{}
    if err := common.UnmarshalBodyReusable(r.c, r.geminiRequest); err != nil {
        return err
    }
    r.geminiRequest.Model = modelList[0]
    r.geminiRequest.Stream = isStream
    r.setOriginalModel(r.geminiRequest.Model)
    r.action = action

    // 检测是否包含音频输入（用于 2.5 Flash 动态切换音频单价）
    hasAudio := false
    if r.geminiRequest != nil {
        for _, content := range r.geminiRequest.Contents {
            for _, part := range content.Parts {
                if part.InlineData != nil {
                    mt := strings.ToLower(strings.TrimSpace(part.InlineData.MimeType))
                    if strings.HasPrefix(mt, "audio/") {
                        hasAudio = true
                        break
                    }
                }
            }
            if hasAudio { break }
        }
    }
    if hasAudio {
        r.c.Set("gemini_audio_input", true)
    }

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
    // 如果 action 为聊天（generateContent/streamGenerateContent），按聊天路径处理；
    // 否则（如 predict/predictLongRunning），透传原始请求至对应 action。
    if r.action != "generateContent" && r.action != "streamGenerateContent" {
        if r.heartbeat != nil {
            r.heartbeat.Stop()
        }
        // 仅支持非流式透传（Imagen/Veo 初始化任务均为 JSON 返回）
        if gp, ok := r.provider.(*gemini.GeminiProvider); ok {
            data, e := gp.RelayModelAction(r.modelName, r.action)
            if e != nil {
                err = e
                done = true
                return
            }
            err = responseJsonClient(r.c, data)
            if err != nil { done = true }
            return
        }
        err = common.StringErrorWrapperLocal("channel not implemented", "channel_error", http.StatusServiceUnavailable)
        done = true
        return
    }

    // 聊天：优先品牌原生Provider
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
        if r.heartbeat != nil { r.heartbeat.Stop() }
        gResp := gemini.OpenAIToGeminiChatResponse(resp, r.modelName, r.provider.GetUsage())
        err = responseJsonClient(r.c, gResp)
        if err != nil { done = true }
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

// GeminiOperations 代理：转发 /gemini/:version/operations/*name 到 Google 原生接口
// 支持使用 query 参数 "model" 指定用于选择渠道的模型名（缺省使用 gemini-2.5-pro）
func GeminiOperations(c *gin.Context) {
    // 允许的渠道类型（与 Gemini 官方路由一致）
    c.Set("allow_channel_type", AllowGeminiChannelType)

    version := strings.TrimSpace(c.Param("version"))
    name := strings.TrimPrefix(c.Param("name"), "/") // operations/...
    if version == "" || name == "" {
        common.AbortWithMessage(c, http.StatusBadRequest, "invalid operations path")
        return
    }

    modelName := strings.TrimSpace(c.Query("model"))
    if modelName == "" {
        modelName = "gemini-2.5-pro"
    }

    provider, _, fail := GetProvider(c, modelName)
    if fail != nil || provider == nil {
        common.AbortWithMessage(c, http.StatusServiceUnavailable, "channel not implemented")
        return
    }

    // 构造原生 URL 与请求头；若上游为 sutui，则改为查询 sutui /v1/videos 并转换为 operations
    if gp, ok := provider.(*gemini.GeminiProvider); ok {
        // sutui 分支
        if gp != nil && strings.EqualFold(gp.DetectVeoVendorForOps(), "sutui") {
            // 从 name 中取出 id（允许传入 operations/<id> 或直接 <id>）
            id := name
            if strings.HasPrefix(id, "operations/") { id = strings.TrimPrefix(id, "operations/") }

            base := strings.TrimSuffix(gp.GetBaseURL(), "/")
            fullURL := fmt.Sprintf("%s/v1/videos/%s", base, id)
            headers := gp.GetRequestHeaders()

            req, err := gp.Requester.NewRequest(http.MethodGet, fullURL, gp.Requester.WithHeader(headers))
            if err != nil {
                common.AbortWithMessage(c, http.StatusInternalServerError, "new_request_failed")
                return
            }
            if req.Body != nil { defer req.Body.Close() }

            var s struct {
                ID        string `json:"id"`
                Status    string `json:"status"`
                Seconds   int    `json:"seconds"`
                Size      string `json:"size"`
                VideoURL  string `json:"video_url"`
                Result    *struct{ VideoURL string `json:"video_url"` } `json:"result"`
            }
            if _, e := gp.Requester.SendRequest(req, &s, false); e != nil {
                gemErr := gemini.OpenaiErrToGeminiErr(e)
                status := e.StatusCode
                if status == 0 { status = http.StatusBadRequest }
                c.JSON(status, gemErr.GeminiErrorResponse)
                return
            }
            // done 判定
            st := strings.ToLower(strings.TrimSpace(s.Status))
            done := (st == "completed" || st == "failed")
            // 取 videoURL
            video := s.VideoURL
            if video == "" && s.Result != nil { video = s.Result.VideoURL }

            // 组装 operations 响应（与官方文档对齐：generatedSamples[].video.uri）
            sample := map[string]any{}
            if video != "" { sample["video"] = map[string]any{"uri": video} }
            meta := map[string]any{}
            if s.Seconds > 0 { meta["durationSeconds"] = s.Seconds }
            if strings.TrimSpace(s.Size) != "" { meta["size"] = s.Size }
            if len(meta) > 0 { sample["metadata"] = meta }

            var samples []any
            if len(sample) > 0 { samples = append(samples, sample) }

            resp := map[string]any{
                "name": name,
                "done": done,
            }
            if done {
                resp["response"] = map[string]any{
                    "generateVideoResponse": map[string]any{
                        "generatedSamples": samples,
                    },
                }
            }

            // 计费：尽量按秒写入
            seconds := s.Seconds
            if seconds == 0 {
                if v := strings.TrimSpace(c.Query("duration")); v != "" {
                    if n, err := strconv.Atoi(v); err == nil { seconds = n }
                }
            }
            if seconds > 0 {
                usage := &types.Usage{PromptTokens: seconds}
                q := relay_util.NewQuota(c, modelName, seconds)
                q.Consume(c, usage, false)
            }

            _ = responseJsonClient(c, resp)
            return
        }

        // Google 原生分支
        base := strings.TrimSuffix(gp.GetBaseURL(), "/")
        fullURL := fmt.Sprintf("%s/%s/%s", base, version, name)
        headers := gp.GetRequestHeaders()

        req, err := gp.Requester.NewRequest(http.MethodGet, fullURL, gp.Requester.WithHeader(headers))
        if err != nil {
            common.AbortWithMessage(c, http.StatusInternalServerError, "new_request_failed")
            return
        }
        defer req.Body.Close()

        var data any
        if _, e := gp.Requester.SendRequest(req, &data, false); e != nil {
            // 统一返回 Gemini 风格的错误
            gemErr := gemini.OpenaiErrToGeminiErr(e)
            status := e.StatusCode
            if status == 0 {
                status = http.StatusBadRequest
            }
            c.JSON(status, gemErr.GeminiErrorResponse)
            return
        }
        // 若提供了 duration（秒）或可从返回中解析，则按秒计费
        seconds := 0
        if durStr := strings.TrimSpace(c.Query("duration")); durStr != "" {
            if n, err := strconv.Atoi(durStr); err == nil && n > 0 { seconds = n }
        }
        // 简单尝试从返回结构提取秒数（若包含）
        if seconds == 0 {
            if m, ok := data.(map[string]any); ok {
                if done, ok2 := m["done"].(bool); ok2 && done {
                    if resp, ok3 := m["response"].(map[string]any); ok3 {
                        if gvr, ok4 := resp["generateVideoResponse"].(map[string]any); ok4 {
                            if samples, ok5 := gvr["generatedSamples"].([]any); ok5 && len(samples) > 0 {
                                if sample, ok6 := samples[0].(map[string]any); ok6 {
                                    if meta, ok7 := sample["metadata"].(map[string]any); ok7 {
                                        if dur, ok8 := meta["durationSeconds"].(float64); ok8 && dur > 0 {
                                            seconds = int(dur)
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }

        if seconds > 0 {
            usage := &types.Usage{PromptTokens: seconds}
            q := relay_util.NewQuota(c, modelName, seconds)
            q.Consume(c, usage, false)
        }

        _ = responseJsonClient(c, data)
        return
    }

    common.AbortWithMessage(c, http.StatusServiceUnavailable, "channel not implemented")
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
