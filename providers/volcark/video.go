package volcark

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"one-api/common/requester"
	"one-api/types"
)

const (
	volcArkVideoTaskPath      = "/api/v3/contents/generations/tasks"
	volcArkVideoTaskQueryPath = "/api/v3/contents/generations/tasks/%s"
)

// VolcArkVideoCreateRequest 使用 map[string]any 透传，单独声明结构方便文档化。
type VolcArkVideoCreateRequest map[string]any

type VolcArkVideoCreateResponse struct {
	ID string `json:"id"`
}

type VolcArkVideoTask struct {
	ID              string                 `json:"id"`
	Model           string                 `json:"model,omitempty"`
	Status          string                 `json:"status,omitempty"`
	Error           *VolcArkVideoError     `json:"error,omitempty"`
	CreatedAt       int64                  `json:"created_at,omitempty"`
	UpdatedAt       int64                  `json:"updated_at,omitempty"`
	Content         *VolcArkVideoTaskMedia `json:"content,omitempty"`
	Seed            *int                   `json:"seed,omitempty"`
	Resolution      string                 `json:"resolution,omitempty"`
	Ratio           string                 `json:"ratio,omitempty"`
	Duration        *int                   `json:"duration,omitempty"`
	FramesPerSecond *int                   `json:"framespersecond,omitempty"`
	Usage           *VolcArkVideoUsage     `json:"usage,omitempty"`
}

type VolcArkVideoTaskMedia struct {
	VideoURL     string `json:"video_url,omitempty"`
	LastFrameURL string `json:"last_frame_url,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
}

type VolcArkVideoUsage struct {
	PromptTokens     int64 `json:"prompt_tokens,omitempty"`
	CompletionTokens int64 `json:"completion_tokens,omitempty"`
	TotalTokens      int64 `json:"total_tokens,omitempty"`
	VideoTokens      int64 `json:"video_tokens,omitempty"`
}

type VolcArkVideoListResponse struct {
	Items []VolcArkVideoTask `json:"items"`
	Total int                `json:"total"`
}

type VolcArkVideoError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateVideoTask 调用 Ark 创建视频生成任务接口。
func (p *VolcArkProvider) CreateVideoTask(body map[string]any) (*VolcArkVideoCreateResponse, *types.OpenAIError) {
	url := p.GetFullRequestURL(volcArkVideoTaskPath, "")
	headers := p.GetRequestHeaders()

	requester := p.newVideoRequester()
	req, err := requester.NewRequest(http.MethodPost, url, requester.WithBody(body), requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{
			Message: fmt.Sprintf("create request failed: %s", err.Error()),
			Type:    "volcark_request_error",
		}
	}

	var resp VolcArkVideoCreateResponse
	_, errWithCode := requester.SendRequest(req, &resp, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &resp, nil
}

// GetVideoTask 查询单个视频任务。
func (p *VolcArkProvider) GetVideoTask(taskID string) (*VolcArkVideoTask, *types.OpenAIError) {
	path := fmt.Sprintf(volcArkVideoTaskQueryPath, url.PathEscape(taskID))
	url := p.GetFullRequestURL(path, "")
	headers := p.GetRequestHeaders()

	requester := p.newVideoRequester()
	req, err := requester.NewRequest(http.MethodGet, url, requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{
			Message: fmt.Sprintf("create request failed: %s", err.Error()),
			Type:    "volcark_request_error",
		}
	}

	var resp VolcArkVideoTask
	_, errWithCode := requester.SendRequest(req, &resp, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &resp, nil
}

// ListVideoTasks 批量查询视频任务。
func (p *VolcArkProvider) ListVideoTasks(params url.Values) (*VolcArkVideoListResponse, *types.OpenAIError) {
	requestURL := volcArkVideoTaskPath
	if len(params) > 0 {
		requestURL = fmt.Sprintf("%s?%s", requestURL, params.Encode())
	}

	url := p.GetFullRequestURL(requestURL, "")
	headers := p.GetRequestHeaders()

	requester := p.newVideoRequester()
	req, err := requester.NewRequest(http.MethodGet, url, requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{
			Message: fmt.Sprintf("create request failed: %s", err.Error()),
			Type:    "volcark_request_error",
		}
	}

	var resp VolcArkVideoListResponse
	_, errWithCode := requester.SendRequest(req, &resp, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &resp, nil
}

// CancelVideoTask 取消或删除视频任务。
func (p *VolcArkProvider) CancelVideoTask(taskID string) *types.OpenAIError {
	path := fmt.Sprintf(volcArkVideoTaskQueryPath, url.PathEscape(taskID))
	url := p.GetFullRequestURL(path, "")
	headers := p.GetRequestHeaders()

	requester := p.newVideoRequester()
	req, err := requester.NewRequest(http.MethodDelete, url, requester.WithHeader(headers))
	if err != nil {
		return &types.OpenAIError{
			Message: fmt.Sprintf("create request failed: %s", err.Error()),
			Type:    "volcark_request_error",
		}
	}

	_, errWithCode := requester.SendRequest(req, nil, false)
	if errWithCode != nil {
		return &errWithCode.OpenAIError
	}

	return nil
}

// RequestErrorHandle 针对视频任务接口的错误解析。
func VideoTaskErrorHandle(resp *http.Response) *types.OpenAIError {
	defer resp.Body.Close()
	var volcErr struct {
		Code    string             `json:"code"`
		Message string             `json:"message"`
		Error   *VolcArkVideoError `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&volcErr); err != nil {
		return &types.OpenAIError{
			Message: fmt.Sprintf("upstream error: %s", resp.Status),
			Type:    "volcark_upstream_error",
			Code:    fmt.Sprintf("%d", resp.StatusCode),
		}
	}

	errPayload := volcErr.Error
	if errPayload == nil {
		errPayload = &VolcArkVideoError{Code: volcErr.Code, Message: volcErr.Message}
	}

	return &types.OpenAIError{
		Message: errPayload.Message,
		Type:    "volcark_upstream_error",
		Code:    errPayload.Code,
	}
}

func (p *VolcArkProvider) newVideoRequester() *requester.HTTPRequester {
	proxy := ""
	if p.Channel != nil && p.Channel.Proxy != nil {
		proxy = *p.Channel.Proxy
	}
	req := requester.NewHTTPRequester(proxy, VideoTaskErrorHandle)
	req.IsOpenAI = false
	req.Context = p.Requester.Context
	return req
}
