package gemini

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"one-api/common"
	"one-api/common/image"
	"one-api/common/storage"
	"one-api/common/utils"
	"one-api/types"
	"regexp"
	"strings"

	goahocorasick "github.com/anknown/ahocorasick"
)

// 将Gemini角色转换为OpenAI角色
func convertGeminiRoleToOpenAI(role string) string {
	switch role {
	case "model":
		return types.ChatMessageRoleAssistant
	case "function":
		return types.ChatMessageRoleTool
	default:
		return types.ChatMessageRoleUser
	}
}

// 将Gemini请求转换为OpenAI ChatCompletion请求
func GeminiToOpenAIChatRequest(req *GeminiChatRequest) (*types.ChatCompletionRequest, *types.OpenAIErrorWithStatusCode) {
	if req == nil {
		return nil, common.StringErrorWrapperLocal("empty request", "invalid_request", http.StatusBadRequest)
	}

	openaiReq := &types.ChatCompletionRequest{
		Model:   req.Model,
		Stream:  req.Stream,
		Messages: make([]types.ChatCompletionMessage, 0),
	}

	// SystemInstruction
	switch v := req.SystemInstruction.(type) {
	case *GeminiChatContent:
		var buf strings.Builder
		for _, p := range v.Parts {
			if p.Text != "" {
				if buf.Len() > 0 { buf.WriteString("\n") }
				buf.WriteString(p.Text)
			}
		}
		if buf.Len() > 0 {
			openaiReq.Messages = append(openaiReq.Messages, types.ChatCompletionMessage{
				Role:    types.ChatMessageRoleSystem,
				Content: buf.String(),
			})
		}
	case GeminiChatContent:
		var buf strings.Builder
		for _, p := range v.Parts {
			if p.Text != "" {
				if buf.Len() > 0 { buf.WriteString("\n") }
				buf.WriteString(p.Text)
			}
		}
		if buf.Len() > 0 {
			openaiReq.Messages = append(openaiReq.Messages, types.ChatCompletionMessage{
				Role:    types.ChatMessageRoleSystem,
				Content: buf.String(),
			})
		}
	}

	for _, c := range req.Contents {
		msg := types.ChatCompletionMessage{Role: convertGeminiRoleToOpenAI(c.Role)}
		parts := make([]map[string]any, 0)

		for _, p := range c.Parts {
			if p.FunctionCall != nil {
				args, _ := json.Marshal(p.FunctionCall.Args)
				msg.ToolCalls = append(msg.ToolCalls, &types.ChatCompletionToolCalls{
					Type: types.ChatMessageRoleFunction,
					Function: &types.ChatCompletionToolCallsFunction{
						Name:      p.FunctionCall.Name,
						Arguments: string(args),
					},
				})
				continue
			}
			if p.FunctionResponse != nil {
				name := p.FunctionResponse.Name
				msg.Role = types.ChatMessageRoleTool
				msg.Name = &name
				// 兼容多种响应结构：优先取 response.content 字段，否则整体序列化
				var contentStr string
				switch v := p.FunctionResponse.Response.(type) {
				case string:
					contentStr = v
				case map[string]any:
					if s, ok := v["content"].(string); ok {
						contentStr = s
					} else {
						b, _ := json.Marshal(v)
						contentStr = string(b)
					}
				default:
					b, _ := json.Marshal(v)
					contentStr = string(b)
				}
				msg.Content = contentStr
				continue
			}
			if p.Text != "" {
				parts = append(parts, map[string]any{"type":"text","text": p.Text})
			}
			if p.InlineData != nil {
				// 使用 data URL 形式传递给OpenAI兼容上游
				url := fmt.Sprintf("data:%s;base64,%s", p.InlineData.MimeType, p.InlineData.Data)
				parts = append(parts, map[string]any{
					"type":"image_url",
					"image_url": map[string]any{"url": url},
				})
			}
			if p.FileData != nil {
				// 简化处理：将fileUri作为文本附加
				parts = append(parts, map[string]any{"type":"text","text": p.FileData.FileUri})
			}
		}

		if len(parts) > 0 && msg.Content == nil {
			msg.Content = parts
		}
		openaiReq.Messages = append(openaiReq.Messages, msg)
	}

	return openaiReq, nil
}

// 将OpenAI响应转换为Gemini响应（非流式）
func OpenAIToGeminiChatResponse(resp *types.ChatCompletionResponse, modelName string, usage *types.Usage) *GeminiChatResponse {
	g := &GeminiChatResponse{
		Model:     modelName,
		ResponseId: resp.ID,
	}
	cands := make([]GeminiChatCandidate, 0, len(resp.Choices))
	for _, ch := range resp.Choices {
		text := ch.Message.StringContent()
		parts := []GeminiPart{}
		if text != "" {
			parts = append(parts, GeminiPart{Text: text})
		}
		if len(ch.Message.Image) > 0 {
			for _, img := range ch.Message.Image {
				parts = append(parts, GeminiPart{InlineData: &GeminiInlineData{MimeType: "image/png", Data: img.Data}})
			}
		}
		cands = append(cands, GeminiChatCandidate{
			Index: int64(ch.Index),
			Content: GeminiChatContent{Role: "model", Parts: parts},
		})
	}
	g.Candidates = cands
	if usage != nil {
		g.UsageMetadata = &GeminiUsageMetadata{
			PromptTokenCount:     usage.PromptTokens,
			CandidatesTokenCount: usage.CompletionTokens,
			TotalTokenCount:      usage.TotalTokens,
		}
	}
	return g
}

const GeminiImageSymbol = "![one-hub-gemini-image]"

const (
	ModalityTEXT  = "TEXT"
	ModalityAUDIO = "AUDIO"
	ModalityIMAGE = "IMAGE"
	ModalityVIDEO = "VIDEO"
)

var ImageSymbolAcMachines = &goahocorasick.Machine{}
var imageRegex = regexp.MustCompile(`\!\[one-hub-gemini-image\]\((.*?)\)`)

func init() {
	ImageSymbolAcMachines.Build([][]rune{[]rune(GeminiImageSymbol)})
}

type GeminiChatRequest struct {
	Model             string                     `json:"-"`
	Stream            bool                       `json:"-"`
	Contents          []GeminiChatContent        `json:"contents"`
	SafetySettings    []GeminiChatSafetySettings `json:"safetySettings,omitempty"`
	GenerationConfig  GeminiChatGenerationConfig `json:"generationConfig,omitempty"`
	Tools             []GeminiChatTools          `json:"tools,omitempty"`
	ToolConfig        *GeminiToolConfig          `json:"toolConfig,omitempty"`
	SystemInstruction any                        `json:"systemInstruction,omitempty"`

	JsonRaw []byte `json:"-"`
}

type GeminiToolConfig struct {
	FunctionCallingConfig *GeminiFunctionCallingConfig `json:"functionCallingConfig,omitempty"`
}

type GeminiFunctionCallingConfig struct {
	Model                string `json:"model,omitempty"`
	AllowedFunctionNames any    `json:"allowedFunctionNames,omitempty"`
}
type GeminiInlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}

type GeminiFileData struct {
	MimeType string `json:"mimeType,omitempty"`
	FileUri  string `json:"fileUri,omitempty"`
}

type GeminiPart struct {
	FunctionCall        *GeminiFunctionCall            `json:"functionCall,omitempty"`
	FunctionResponse    *GeminiFunctionResponse        `json:"functionResponse,omitempty"`
	Text                string                         `json:"text,omitempty"`
	InlineData          *GeminiInlineData              `json:"inlineData,omitempty"`
	FileData            *GeminiFileData                `json:"fileData,omitempty"`
	ExecutableCode      *GeminiPartExecutableCode      `json:"executableCode,omitempty"`
	CodeExecutionResult *GeminiPartCodeExecutionResult `json:"codeExecutionResult,omitempty"`
	Thought             bool                           `json:"thought,omitempty"` // 是否是思考内容
}

type GeminiPartExecutableCode struct {
	Language string `json:"language,omitempty"`
	Code     string `json:"code,omitempty"`
}

type GeminiPartCodeExecutionResult struct {
	Outcome string `json:"outcome,omitempty"`
	Output  string `json:"output,omitempty"`
}

type GeminiFunctionCall struct {
	Name string                 `json:"name,omitempty"`
	Args map[string]interface{} `json:"args,omitempty"`
}

func (candidate *GeminiChatCandidate) ToOpenAIStreamChoice(request *types.ChatCompletionRequest) types.ChatCompletionStreamChoice {
	choice := types.ChatCompletionStreamChoice{
		Index: int(candidate.Index),
		Delta: types.ChatCompletionStreamChoiceDelta{
			Role: types.ChatMessageRoleAssistant,
		},
	}

	if candidate.FinishReason != nil {
		choice.FinishReason = ConvertFinishReason(*candidate.FinishReason)
	}

	var content []string
	isTools := false
	images := make([]types.MultimediaData, 0)
	reasoningContent := make([]string, 0)

	for _, part := range candidate.Content.Parts {
		if part.FunctionCall != nil {
			if choice.Delta.ToolCalls == nil {
				choice.Delta.ToolCalls = make([]*types.ChatCompletionToolCalls, 0)
			}
			isTools = true
			choice.Delta.ToolCalls = append(choice.Delta.ToolCalls, part.FunctionCall.ToOpenAITool())
		} else if part.InlineData != nil {
			if strings.HasPrefix(part.InlineData.MimeType, "image/") {
				images = append(images, types.MultimediaData{
					Data: part.InlineData.Data,
				})
				url := ""
				imageData, err := base64.StdEncoding.DecodeString(part.InlineData.Data)
				if err == nil {
					url = storage.Upload(imageData, utils.GetUUID()+".png")
				}
				if url == "" {
					url = "image upload err"
				}
				content = append(content, fmt.Sprintf("%s(%s)", GeminiImageSymbol, url))
			}
			//  else if strings.HasPrefix(part.InlineData.MimeType, "audio/") {
			// 	choice.Message.Audio = types.MultimediaData{
			// 		Data: part.InlineData.Data,
			// 	}
			// }
		} else {
			if part.ExecutableCode != nil {
				content = append(content, "```"+part.ExecutableCode.Language+"\n"+part.ExecutableCode.Code+"\n```")
			} else if part.CodeExecutionResult != nil {
				content = append(content, "```output\n"+part.CodeExecutionResult.Output+"\n```")
			} else if part.Thought {
				reasoningContent = append(reasoningContent, part.Text)
			} else {
				content = append(content, part.Text)
			}
		}
	}

	if len(images) > 0 {
		choice.Delta.Image = images
	}

	choice.Delta.Content = strings.Join(content, "\n")

	if len(reasoningContent) > 0 {
		choice.Delta.ReasoningContent = strings.Join(reasoningContent, "\n")
	}

	if isTools {
		choice.FinishReason = types.FinishReasonToolCalls
	}
	choice.CheckChoice(request)

	return choice
}

func (candidate *GeminiChatCandidate) ToOpenAIChoice(request *types.ChatCompletionRequest) types.ChatCompletionChoice {
	choice := types.ChatCompletionChoice{
		Index: int(candidate.Index),
		Message: types.ChatCompletionMessage{
			Role: "assistant",
		},
		// FinishReason: types.FinishReasonStop,
	}

	if candidate.FinishReason != nil {
		choice.FinishReason = ConvertFinishReason(*candidate.FinishReason)
	}

	if len(candidate.Content.Parts) == 0 {
		choice.Message.Content = ""
		return choice
	}

	var content []string
	useTools := false
	images := make([]types.MultimediaData, 0)
	reasoningContent := make([]string, 0)

	for _, part := range candidate.Content.Parts {
		if part.FunctionCall != nil {
			if choice.Message.ToolCalls == nil {
				choice.Message.ToolCalls = make([]*types.ChatCompletionToolCalls, 0)
			}
			useTools = true
			choice.Message.ToolCalls = append(choice.Message.ToolCalls, part.FunctionCall.ToOpenAITool())
		} else if part.InlineData != nil {
			if strings.HasPrefix(part.InlineData.MimeType, "image/") {

				images = append(images, types.MultimediaData{
					Data: part.InlineData.Data,
				})
				url := ""
				imageData, err := base64.StdEncoding.DecodeString(part.InlineData.Data)
				if err == nil {
					url = storage.Upload(imageData, utils.GetUUID()+".png")
				}
				if url == "" {
					url = "image upload err"
				}
				content = append(content, fmt.Sprintf("%s(%s)", GeminiImageSymbol, url))
			}
			//  else if strings.HasPrefix(part.InlineData.MimeType, "audio/") {
			// 	choice.Message.Audio = types.MultimediaData{
			// 		Data: part.InlineData.Data,
			// 	}
			// }
		} else {
			if part.ExecutableCode != nil {
				content = append(content, "```"+part.ExecutableCode.Language+"\n"+part.ExecutableCode.Code+"\n```")
			} else if part.CodeExecutionResult != nil {
				content = append(content, "```output\n"+part.CodeExecutionResult.Output+"\n```")
			} else if part.Thought {
				reasoningContent = append(reasoningContent, part.Text)
			} else {
				content = append(content, part.Text)
			}
		}
	}

	choice.Message.Content = strings.Join(content, "\n")

	if len(reasoningContent) > 0 {
		choice.Message.ReasoningContent = strings.Join(reasoningContent, "\n")
	}

	if len(images) > 0 {
		choice.Message.Image = images
	}

	if useTools {
		choice.FinishReason = types.FinishReasonToolCalls
	}

	choice.CheckChoice(request)

	return choice
}

type GeminiFunctionResponse struct {
	Name     string `json:"name,omitempty"`
	Response any    `json:"response,omitempty"`
}

type GeminiFunctionResponseContent struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

func (g *GeminiFunctionCall) ToOpenAITool() *types.ChatCompletionToolCalls {
	args, _ := json.Marshal(g.Args)

	return &types.ChatCompletionToolCalls{
		Id:    "call_" + utils.GetRandomString(24),
		Type:  types.ChatMessageRoleFunction,
		Index: 0,
		Function: &types.ChatCompletionToolCallsFunction{
			Name:      g.Name,
			Arguments: string(args),
		},
	}
}

type GeminiChatContent struct {
	Role  string       `json:"role,omitempty"`
	Parts []GeminiPart `json:"parts,omitempty"`
}

type GeminiChatSafetySettings struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

type GeminiChatTools struct {
	FunctionDeclarations  []types.ChatCompletionFunction `json:"functionDeclarations,omitempty"`
	CodeExecution         *GeminiCodeExecution           `json:"codeExecution,omitempty"`
	GoogleSearch          any                            `json:"googleSearch,omitempty"`
	UrlContext            any                            `json:"urlContext,omitempty"`
	GoogleSearchRetrieval any                            `json:"googleSearchRetrieval,omitempty"`
}

type GeminiCodeExecution struct {
}

type GeminiChatGenerationConfig struct {
	Temperature        *float64        `json:"temperature,omitempty"`
	TopP               *float64        `json:"topP,omitempty"`
	TopK               *float64        `json:"topK,omitempty"`
	MaxOutputTokens    int             `json:"maxOutputTokens,omitempty"`
	CandidateCount     int             `json:"candidateCount,omitempty"`
	StopSequences      []string        `json:"stopSequences,omitempty"`
	ResponseMimeType   string          `json:"responseMimeType,omitempty"`
	ResponseSchema     any             `json:"responseSchema,omitempty"`
	ResponseModalities []string        `json:"responseModalities,omitempty"`
	ThinkingConfig     *ThinkingConfig `json:"thinkingConfig,omitempty"`
}

type ThinkingConfig struct {
	ThinkingBudget  *int `json:"thinkingBudget"`
	IncludeThoughts bool `json:"includeThoughts,omitempty"`
}

type GeminiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func (e *GeminiError) Error() string {
	bytes, _ := json.Marshal(e)
	return string(bytes) + "\n"
}

type GeminiErrorResponse struct {
	ErrorInfo *GeminiError `json:"error,omitempty"`
}

func (e *GeminiErrorResponse) Error() string {
	bytes, _ := json.Marshal(e)
	return string(bytes) + "\n"
}

type GeminiChatResponse struct {
	Candidates     []GeminiChatCandidate    `json:"candidates"`
	PromptFeedback GeminiChatPromptFeedback `json:"promptFeedback"`
	UsageMetadata  *GeminiUsageMetadata     `json:"usageMetadata,omitempty"`
	ModelVersion   string                   `json:"modelVersion,omitempty"`
	Model          string                   `json:"model,omitempty"`
	ResponseId     string                   `json:"responseId,omitempty"`
	GeminiErrorResponse
}

type GeminiUsageMetadata struct {
	PromptTokenCount        int `json:"promptTokenCount"`
	CandidatesTokenCount    int `json:"candidatesTokenCount"`
	TotalTokenCount         int `json:"totalTokenCount"`
	CachedContentTokenCount int `json:"cachedContentTokenCount,omitempty"`
	ThoughtsTokenCount      int `json:"thoughtsTokenCount,omitempty"`
	ToolUsePromptTokenCount int `json:"toolUsePromptTokenCount,omitempty"`

	PromptTokensDetails     []GeminiUsageMetadataDetails `json:"promptTokensDetails,omitempty"`
	CandidatesTokensDetails []GeminiUsageMetadataDetails `json:"candidatesTokensDetails,omitempty"`
}

type GeminiUsageMetadataDetails struct {
	Modality   string `json:"modality"`
	TokenCount int    `json:"tokenCount"`
}

type GeminiChatCandidate struct {
	Content               GeminiChatContent        `json:"content"`
	FinishReason          *string                  `json:"finishReason,omitempty"`
	Index                 int64                    `json:"index"`
	SafetyRatings         []GeminiChatSafetyRating `json:"safetyRatings"`
	CitationMetadata      any                      `json:"citationMetadata,omitempty"`
	TokenCount            int                      `json:"tokenCount,omitempty"`
	GroundingAttributions []any                    `json:"groundingAttributions,omitempty"`
	AvgLogprobs           any                      `json:"avgLogprobs,omitempty"`
}

type GeminiChatSafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type GeminiChatPromptFeedback struct {
	BlockReason   string                   `json:"blockReason"`
	SafetyRatings []GeminiChatSafetyRating `json:"safetyRatings"`
}

func (g *GeminiChatResponse) GetResponseText() string {
	if g == nil {
		return ""
	}
	if len(g.Candidates) > 0 && len(g.Candidates[0].Content.Parts) > 0 {
		return g.Candidates[0].Content.Parts[0].Text
	}
	return ""
}

func OpenAIToGeminiChatContent(openaiContents []types.ChatCompletionMessage) ([]GeminiChatContent, string, *types.OpenAIErrorWithStatusCode) {
	contents := make([]GeminiChatContent, 0)
	// useToolName := ""
	var systemContent []string
	toolCallId := make(map[string]string)

	for _, openaiContent := range openaiContents {
		if openaiContent.IsSystemRole() {
			systemContent = append(systemContent, openaiContent.StringContent())
			continue
		}

		content := GeminiChatContent{
			Role:  ConvertRole(openaiContent.Role),
			Parts: make([]GeminiPart, 0),
		}
		openaiContent.FuncToToolCalls()

		if openaiContent.ToolCalls != nil {
			for _, toolCall := range openaiContent.ToolCalls {
				toolCallId[toolCall.Id] = toolCall.Function.Name

				args := map[string]interface{}{}
				if toolCall.Function.Arguments != "" {
					json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
				}

				content.Parts = append(content.Parts, GeminiPart{
					FunctionCall: &GeminiFunctionCall{
						Name: toolCall.Function.Name,
						Args: args,
					},
				})

			}
			text := openaiContent.StringContent()
			if text != "" {
				contents = append(contents, createSystemResponse(text))
			}
		} else if openaiContent.Role == types.ChatMessageRoleFunction || openaiContent.Role == types.ChatMessageRoleTool {
			if openaiContent.Name == nil {
				if toolName, exists := toolCallId[openaiContent.ToolCallID]; exists {
					openaiContent.Name = &toolName
				}
			}

			functionPart := GeminiPart{
				FunctionResponse: &GeminiFunctionResponse{
					Name: *openaiContent.Name,
					Response: GeminiFunctionResponseContent{
						Name:    *openaiContent.Name,
						Content: openaiContent.StringContent(),
					},
				},
			}

			if len(contents) > 0 && contents[len(contents)-1].Role == "function" {
				contents[len(contents)-1].Parts = append(contents[len(contents)-1].Parts, functionPart)
			} else {
				contents = append(contents, GeminiChatContent{
					Role:  "function",
					Parts: []GeminiPart{functionPart},
				})
			}

			continue
		} else {
			openaiMessagePart := openaiContent.ParseContent()
			imageNum := 0
			for _, openaiPart := range openaiMessagePart {
				if openaiPart.Type == types.ContentTypeText {
					if openaiPart.Text == "" {
						continue
					}
					imageSymbols := ImageSymbolAcMachines.MultiPatternSearch([]rune(openaiPart.Text), false)
					if len(imageSymbols) > 0 {
						lastEndPos := 0 // 上一段文本的结束位置
						textRunes := []rune(openaiPart.Text)
						geminiImageSymbolRunesLen := len([]rune(GeminiImageSymbol))
						// 提取图片地址
						for _, match := range imageSymbols {
							// 添加图片符号前面的文本，如果不为空且不仅包含换行符
							if match.Pos > lastEndPos {
								textSegment := string(textRunes[lastEndPos:match.Pos])
								if !isEmptyOrOnlyNewlines(textSegment) {
									content.Parts = append(content.Parts, GeminiPart{
										Text: textSegment,
									})
								}
							}

							pos := match.Pos + geminiImageSymbolRunesLen

							if pos < len(textRunes) && textRunes[pos] == '(' {
								endPos := -1
								for i := pos + 1; i < len(textRunes); i++ {
									if textRunes[i] == ')' {
										endPos = i
										break
									}
								}
								if endPos > 0 {
									imageUrl := string(textRunes[pos+1 : endPos])
									// 处理图片URL
									mimeType, data, err := image.GetImageFromUrl(imageUrl)
									if err == nil {
										content.Parts = append(content.Parts, GeminiPart{
											InlineData: &GeminiInlineData{
												MimeType: mimeType,
												Data:     data,
											},
										})
									}
									lastEndPos = endPos + 1
								}
							}

							// 添加最后一个图片符号后面的文本，如果不为空且不仅包含换行符
							if lastEndPos < len(textRunes) {
								finalText := string(textRunes[lastEndPos:])
								if !isEmptyOrOnlyNewlines(finalText) {
									content.Parts = append(content.Parts, GeminiPart{
										Text: finalText,
									})
								}
							}
						}
					} else {
						content.Parts = append(content.Parts, GeminiPart{
							Text: openaiPart.Text,
						})
					}

				} else if openaiPart.Type == types.ContentTypeImageURL {
					imageNum += 1
					if imageNum > GeminiVisionMaxImageNum {
						continue
					}
					mimeType, data, err := image.GetImageFromUrl(openaiPart.ImageURL.URL)
					if err != nil {
						return nil, "", common.ErrorWrapper(err, "image_url_invalid", http.StatusBadRequest)
					}
					content.Parts = append(content.Parts, GeminiPart{
						InlineData: &GeminiInlineData{
							MimeType: mimeType,
							Data:     data,
						},
					})
				}
			}
		}
		contents = append(contents, content)

	}

	return contents, strings.Join(systemContent, "\n"), nil
}

func createSystemResponse(text string) GeminiChatContent {
	return GeminiChatContent{
		Role: "model",
		Parts: []GeminiPart{
			{
				Text: text,
			},
		},
	}
}

type ModelListResponse struct {
	Models []ModelDetails `json:"models"`
}

type ModelDetails struct {
	Name                       string   `json:"name"`
	DisplayName                string   `json:"displayName"`
	SupportedGenerationMethods []string `json:"supportedGenerationMethods"`
}

type GeminiErrorWithStatusCode struct {
	GeminiErrorResponse
	StatusCode int  `json:"status_code"`
	LocalError bool `json:"-"`
}

func (e *GeminiErrorWithStatusCode) ToOpenAiError() *types.OpenAIErrorWithStatusCode {
	return &types.OpenAIErrorWithStatusCode{
		StatusCode: e.StatusCode,
		OpenAIError: types.OpenAIError{
			Code:    e.ErrorInfo.Code,
			Type:    e.ErrorInfo.Status,
			Message: e.ErrorInfo.Message,
		},
		LocalError: e.LocalError,
	}
}

type GeminiErrors []*GeminiErrorResponse

func (e *GeminiErrors) Error() *GeminiErrorResponse {
	return (*e)[0]
}

type GeminiImageRequest struct {
	Instances  []GeminiImageInstance `json:"instances"`
	Parameters GeminiImageParameters `json:"parameters"`
}

type GeminiImageInstance struct {
	Prompt string `json:"prompt"`
}

type GeminiImageParameters struct {
	PersonGeneration string `json:"personGeneration,omitempty"`
	AspectRatio      string `json:"aspectRatio,omitempty"`
	SampleCount      int    `json:"sampleCount,omitempty"`
}

type GeminiImageResponse struct {
	Predictions []GeminiImagePrediction `json:"predictions"`
}

type GeminiImagePrediction struct {
	BytesBase64Encoded string `json:"bytesBase64Encoded"`
	MimeType           string `json:"mimeType"`
	RaiFilteredReason  string `json:"raiFilteredReason,omitempty"`
	SafetyAttributes   any    `json:"safetyAttributes,omitempty"`
}

func isEmptyOrOnlyNewlines(s string) bool {
	trimmed := strings.TrimSpace(s)
	return trimmed == ""
}
