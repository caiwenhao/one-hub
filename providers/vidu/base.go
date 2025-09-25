package vidu

import (
    "encoding/json"
    "fmt"
    "net/http"
    "one-api/common/requester"
    "one-api/model"
    "one-api/providers/base"
    "one-api/types"
    "strings"
)

// 定义供应商工厂
type ViduProviderFactory struct{}

// 创建 ViduProvider
func (f ViduProviderFactory) Create(channel *model.Channel) base.ProviderInterface {
	return &ViduProvider{
		BaseProvider: base.BaseProvider{
			Config:    getConfig(),
			Channel:   channel,
			Requester: requester.NewHTTPRequester(*channel.Proxy, RequestErrorHandle),
		},
		SubmitPath: "/ent/v2/%s",                  // POST
		QueryPath:  "/ent/v2/tasks/%s/creations", // GET
		CancelPath: "/ent/v2/tasks/%s/cancel",    // POST
	}
}

func getConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL: "https://api.vidu.cn",
	}
}

type ViduProvider struct {
	base.BaseProvider
	SubmitPath string
	QueryPath  string
	CancelPath string
}

func (p *ViduProvider) GetRequestHeaders() (headers map[string]string) {
    headers = make(map[string]string)
    p.CommonRequestHeaders(headers)
    if p.Channel.Key != "" {
        headers["Authorization"] = fmt.Sprintf("Token %s", p.Channel.Key)
    }
    headers["Content-Type"] = "application/json"
    return headers
}

// 提交任务 - 对齐官方文档
func (p *ViduProvider) Submit(action string, request interface{}) (*ViduResponse, *types.OpenAIError) {
    url := p.GetFullRequestURL(fmt.Sprintf(p.SubmitPath, action), "")
    headers := p.GetRequestHeaders()

    // 根据不同action类型处理请求
    var reqBody interface{}
    switch action {
    case ViduActionReference2Image:
        reqBody = request
    default:
        // 处理通用任务请求
        if taskReq, ok := request.(*ViduTaskRequest); ok {
            // 规范化模型名称
            normalizedReq := *taskReq
            normalizedReq.Model = normalizeModelName(taskReq.Model)
            reqBody = &normalizedReq
        } else {
            reqBody = request
        }
    }

    req, err := p.Requester.NewRequest(http.MethodPost, url, p.Requester.WithBody(reqBody), p.Requester.WithHeader(headers))
    if err != nil {
        return nil, &types.OpenAIError{
            Message: fmt.Sprintf("create request failed: %s", err.Error()),
            Type:    "vidu_request_error",
        }
    }

	var response ViduResponse
	_, errWithCode := p.Requester.SendRequest(req, &response, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &response, nil
}

// 查询任务状态
func (p *ViduProvider) QueryCreations(taskId string) (*ViduQueryResponse, *types.OpenAIError) {
	url := p.GetFullRequestURL(fmt.Sprintf(p.QueryPath, taskId), "")
	headers := p.GetRequestHeaders()

	req, err := p.Requester.NewRequest(http.MethodGet, url, p.Requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{
			Message: fmt.Sprintf("create request failed: %s", err.Error()),
			Type:    "vidu_request_error",
		}
	}

	var response ViduQueryResponse
	_, errWithCode := p.Requester.SendRequest(req, &response, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &response, nil
}

// 参考生图接口
func (p *ViduProvider) SubmitReference2Image(request *ViduReference2ImageRequest) (*ViduResponse, *types.OpenAIError) {
	return p.Submit(ViduActionReference2Image, request)
}

// 取消任务
func (p *ViduProvider) CancelTask(taskId string) (*ViduResponse, *types.OpenAIError) {
	url := p.GetFullRequestURL(fmt.Sprintf(p.CancelPath, taskId), "")
	headers := p.GetRequestHeaders()

	cancelReq := ViduCancelRequest{ID: taskId}
	req, err := p.Requester.NewRequest(http.MethodPost, url, p.Requester.WithBody(&cancelReq), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, &types.OpenAIError{
			Message: fmt.Sprintf("create request failed: %s", err.Error()),
			Type:    "vidu_request_error",
		}
	}

	// 取消成功返回空响应，失败返回错误
	var response ViduResponse
	_, errWithCode := p.Requester.SendRequest(req, &response, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &response, nil
}

// 请求错误处理
func RequestErrorHandle(resp *http.Response) *types.OpenAIError {
	errorResponse := &ViduErrorResponse{}
	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return nil
	}

	return ErrorHandle(errorResponse)
}

// 错误处理
func ErrorHandle(err *ViduErrorResponse) *types.OpenAIError {
    if err.Code == 0 {
        return nil
    }

	return &types.OpenAIError{
		Code:    fmt.Sprintf("%d", err.Code),
		Message: err.Message,
		Type:    "vidu_error",
    }
}

// 模型名称规范化 - 对齐官方文档支持的模型
func normalizeModelName(model string) string {
    m := strings.TrimSpace(strings.ToLower(model))
    switch m {
    // 新模型
    case "viduq2-pro":
        return ViduModelQ2Pro
    case "viduq2-turbo":
        return ViduModelQ2Turbo
    // 原有模型
    case "viduq1":
        return ViduModelQ1
    case "viduq1-classic":
        return ViduModelQ1Classic
    case "vidu2.0":
        return ViduModel20
    case "vidu1.5":
        return ViduModel15
    // 向后兼容的别名
    case "vidu-1.5":
        return ViduModel15
    case "vidu-2.0":
        return ViduModel20
    case "vidu-q1":
        return ViduModelQ1
    case "vidu-q1-classic":
        return ViduModelQ1Classic
    case "vidu-q2-pro":
        return ViduModelQ2Pro
    case "vidu-q2-turbo":
        return ViduModelQ2Turbo
    default:
        // 其他输入保持原样，避免阻断新模型接入
        return model
    }
}
