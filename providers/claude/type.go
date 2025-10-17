package claude

import (
	"encoding/json"
	"net/http"
	"one-api/common"
	"one-api/types"
)

const (
	FinishReasonEndTurn = "end_turn"
	FinishReasonToolUse = "tool_use"
)

const (
	ContentTypeText             = "text"
	ContentTypeImage            = "image"
	ContentTypeToolUes          = "tool_use"
	ContentTypeToolResult       = "tool_result"
	ContentTypeThinking         = "thinking"
	ContentTypeRedactedThinking = "redacted_thinking"

	ContentStreamTypeThinking       = "thinking_delta"
	ContentStreamTypeSignatureDelta = "signature_delta"
	ContentStreamTypeInputJsonDelta = "input_json_delta"
)

type ClaudeError struct {
	Type      string          `json:"type"`
	ErrorInfo ClaudeErrorInfo `json:"error"`
}

func (e *ClaudeError) Error() string {
	bytes, _ := json.Marshal(e)
	return string(bytes) + "\n"
}

type ClaudeErrorWithStatusCode struct {
	ClaudeError
	StatusCode int  `json:"status_code"`
	LocalError bool `json:"-"`
}

func (e *ClaudeErrorWithStatusCode) ToOpenAiError() *types.OpenAIErrorWithStatusCode {
	return &types.OpenAIErrorWithStatusCode{
		StatusCode: e.StatusCode,
		OpenAIError: types.OpenAIError{
			Type:    e.Type,
			Message: e.ErrorInfo.Message,
		},
		LocalError: e.LocalError,
	}
}

type ClaudeErrorInfo struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ClaudeMetadata struct {
	UserId string `json:"user_id"`
}

type ResContent struct {
	Text       string `json:"text,omitempty"`
	Type       string `json:"type"`
	Name       string `json:"name,omitempty"`
	Input      any    `json:"input,omitempty"`
	Id         string `json:"id,omitempty"`
	Thinking   string `json:"thinking,omitempty"`
	Signature  string `json:"signature,omitempty"`
	Delta      string `json:"delta,omitempty"`
	Citations  any    `json:"citations,omitempty"`
	Content    any    `json:"content,omitempty"`
	ToolUseId  string `json:"tool_use_id,omitempty"`
	ServerName string `json:"server_name,omitempty"`
	IsError    *bool  `json:"is_error,omitempty"`
	FileId     string `json:"file_id,omitempty"`
	Data       string `json:"data,omitempty"`
}

func (g *ResContent) ToOpenAITool() *types.ChatCompletionToolCalls {
	args, _ := json.Marshal(g.Input)

	return &types.ChatCompletionToolCalls{
		Id:    g.Id,
		Type:  types.ChatMessageRoleFunction,
		Index: 0,
		Function: &types.ChatCompletionToolCallsFunction{
			Name:      g.Name,
			Arguments: string(args),
		},
	}
}

type ContentSource struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type,omitempty"`
	Data      string `json:"data,omitempty"`
	Url       string `json:"url,omitempty"`
}

type MessageContent struct {
	Type         string         `json:"type"`
	Text         string         `json:"text,omitempty"`
	Source       *ContentSource `json:"source,omitempty"`
	Id           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Input        any            `json:"input,omitempty"`
	Content      any            `json:"content,omitempty"`
	IsError      *bool          `json:"is_error,omitempty"`
	ToolUseId    string         `json:"tool_use_id,omitempty"`
	CacheControl any            `json:"cache_control,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content any    `json:"content"`
}

type ClaudeRequest struct {
	Model         string      `json:"model,omitempty"`
	System        any         `json:"system,omitempty"`
	Messages      []Message   `json:"messages"`
	MaxTokens     int         `json:"max_tokens"`
	StopSequences []string    `json:"stop_sequences,omitempty"`
	Temperature   *float64    `json:"temperature,omitempty"`
	TopP          *float64    `json:"top_p,omitempty"`
	TopK          *int        `json:"top_k,omitempty"`
	Tools         []Tools     `json:"tools,omitempty"`
	ToolChoice    *ToolChoice `json:"tool_choice,omitempty"`
	Thinking      *Thinking   `json:"thinking,omitempty"`
	McpServers    any         `json:"mcp_servers,omitempty"`
	//ClaudeMetadata    `json:"metadata,omitempty"`
	Stream bool `json:"stream,omitempty"`
}

type Thinking struct {
	Type         string `json:"type,omitempty"`
	BudgetTokens int    `json:"budget_tokens,omitempty"`
}
type ToolChoice struct {
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	DisableParallelToolUse bool   `json:"disable_parallel_tool_use,omitempty"`
}

type Tools struct {
	Type            string `json:"type,omitempty"`
	CacheControl    any    `json:"cache_control,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	InputSchema     any    `json:"input_schema,omitempty"`
	DisplayHeightPx int    `json:"display_height_px,omitempty"`
	DisplayWidthPx  int    `json:"display_width_px,omitempty"`
	DisplayNumber   int    `json:"display_number,omitempty"`
}

type Usage struct {
	InputTokens              int `json:"input_tokens,omitempty"`
	OutputTokens             int `json:"output_tokens,omitempty"`
	CacheCreationInputTokens int `json:"cache_creation_input_tokens,omitempty"`
	CacheReadInputTokens     int `json:"cache_read_input_tokens,omitempty"`
	CacheCreation            any `json:"cache_creation,omitempty"`

	ServerToolUse *ServerToolUse `json:"server_tool_use,omitempty"`
}

type ServerToolUse struct {
	WebSearchRequests int `json:"web_search_requests,omitempty"`
}
type ClaudeResponse struct {
	Id           string       `json:"id"`
	Type         string       `json:"type"`
	Role         string       `json:"role"`
	Content      []ResContent `json:"content"`
	Model        string       `json:"model"`
	StopReason   string       `json:"stop_reason,omitempty"`
	StopSequence string       `json:"stop_sequence,omitempty"`
	Usage        Usage        `json:"usage,omitempty"`
	Error        *ClaudeError `json:"error,omitempty"`

	Container any `json:"container,omitempty"`
}

type Delta struct {
	Type         string `json:"type,omitempty"`
	Text         string `json:"text,omitempty"`
	PartialJson  string `json:"partial_json,omitempty"`
	StopReason   string `json:"stop_reason,omitempty"`
	StopSequence string `json:"stop_sequence,omitempty"`
	Thinking     string `json:"thinking,omitempty"`
	Signature    string `json:"signature,omitempty"`
	Citations    any    `json:"citations,omitempty"`
}

type ClaudeStreamResponse struct {
	Type         string         `json:"type"`
	Message      ClaudeResponse `json:"message,omitempty"`
	Index        int            `json:"index,omitempty"`
	Delta        Delta          `json:"delta,omitempty"`
	ContentBlock ContentBlock   `json:"content_block,omitempty"`
	Usage        Usage          `json:"usage,omitempty"`
	Error        *ClaudeError   `json:"error,omitempty"`
}

type ContentBlock struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Input any    `json:"input,omitempty"`
	Text  string `json:"text,omitempty"`
}

type ModelListResponse struct {
	Data []Model `json:"data"`
}

type Model struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// 简化：将Claude请求转换为OpenAI ChatCompletion请求（仅文本与基础角色）
func ClaudeToOpenAIChatRequest(req *ClaudeRequest) (*types.ChatCompletionRequest, *types.OpenAIErrorWithStatusCode) {
	if req == nil {
		return nil, common.StringErrorWrapperLocal("empty request", "invalid_request", http.StatusBadRequest)
	}
	oai := &types.ChatCompletionRequest{
		Model:    req.Model,
		Stream:   req.Stream,
		Messages: make([]types.ChatCompletionMessage, 0),
	}
	// system
	switch v := req.System.(type) {
	case string:
		if v != "" {
			oai.Messages = append(oai.Messages, types.ChatCompletionMessage{Role: types.ChatMessageRoleSystem, Content: v})
		}
	case map[string]any:
		if txt, ok := v["text"].(string); ok && txt != "" {
			oai.Messages = append(oai.Messages, types.ChatCompletionMessage{Role: types.ChatMessageRoleSystem, Content: txt})
		}
	}
	for _, m := range req.Messages {
		role := m.Role
		if role == "assistant" { role = types.ChatMessageRoleAssistant }
		if role == "user" { role = types.ChatMessageRoleUser }
		msg := types.ChatCompletionMessage{Role: role}
		switch c := m.Content.(type) {
		case string:
			msg.Content = c
		case []any:
			parts := make([]map[string]any, 0)
			for _, it := range c {
				if mp, ok := it.(map[string]any); ok {
					if t, ok := mp["type"].(string); ok && t == "text" {
						if txt, ok := mp["text"].(string); ok && txt != "" {
							parts = append(parts, map[string]any{"type":"text","text":txt})
						}
					}
				}
			}
			if len(parts) > 0 { msg.Content = parts }
		}
		oai.Messages = append(oai.Messages, msg)
	}
	return oai, nil
}

// 简化：将OpenAI响应转换为Claude响应（仅文本）
func OpenAIToClaudeResponse(resp *types.ChatCompletionResponse, modelName string, usage *types.Usage) *ClaudeResponse {
	content := make([]ResContent, 0)
	var text string
	if resp != nil {
		text = resp.GetContent()
	}
	if text != "" {
		content = append(content, ResContent{Type: ContentTypeText, Text: text})
	}
	cr := &ClaudeResponse{
		Id:     resp.ID,
		Type:   "message",
		Role:   "assistant",
		Content: content,
		Model:  modelName,
		StopReason: FinishReasonEndTurn,
	}
	if usage != nil {
		cr.Usage = Usage{InputTokens: usage.PromptTokens, OutputTokens: usage.CompletionTokens}
	}
	return cr
}
