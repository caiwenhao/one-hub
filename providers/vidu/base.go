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
		SubmitPath: "/ent/v2/%s",              // POST
		QueryPath:  "/ent/v2/tasks/%s/creations", // GET
	}
}

func getConfig() base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL: "https://api.vidu.com",
	}
}

type ViduProvider struct {
	base.BaseProvider
	SubmitPath string
	QueryPath  string
}

func (p *ViduProvider) GetRequestHeaders() (headers map[string]string) {
    headers = make(map[string]string)
    p.CommonRequestHeaders(headers)
    if p.Channel.Key != "" {
        // 认证头可配置：当 Channel.Other 包含 "auth=bearer"（不区分大小写）时，使用 Bearer；否则默认 Token
        authType := "token"
        if strings.Contains(strings.ToLower(p.Channel.Other), "auth=bearer") {
            authType = "bearer"
        }
        if authType == "bearer" {
            headers["Authorization"] = fmt.Sprintf("Bearer %s", p.Channel.Key)
        } else {
            headers["Authorization"] = fmt.Sprintf("Token %s", p.Channel.Key)
        }
    }
    headers["Content-Type"] = "application/json"
    return headers
}

// 提交任务
func (p *ViduProvider) Submit(action string, request *ViduTaskRequest) (*ViduResponse[*ViduTaskData], *types.OpenAIError) {
    url := p.GetFullRequestURL(fmt.Sprintf(p.SubmitPath, action), "")
    headers := p.GetRequestHeaders()

    // 为与官方接口对齐：将模型名规范化为官方API支持的格式（不影响计费所用的 OriginalModel）
    // 根据官方文档，支持的模型格式：viduq1, vidu1.5（保持原格式）
    // 同时支持连字符格式的向后兼容：vidu-1.5 -> vidu1.5
    // 若非上述别名，则保持原样，避免阻断新模型接入。
    reqBody := *request
    if reqBody.Model != "" {
        reqBody.Model = normalizeModelName(reqBody.Model)
    }

    req, err := p.Requester.NewRequest(http.MethodPost, url, p.Requester.WithBody(&reqBody), p.Requester.WithHeader(headers))
    if err != nil {
        return nil, &types.OpenAIError{
            Message: fmt.Sprintf("create request failed: %s", err.Error()),
            Type:    "vidu_request_error",
        }
    }

	var response ViduResponse[*ViduTaskData]
	_, errWithCode := p.Requester.SendRequest(req, &response, false)
	if errWithCode != nil {
		return nil, &errWithCode.OpenAIError
	}

	return &response, nil
}

// 查询任务状态（旧版本接口，保持向后兼容）
func (p *ViduProvider) Query(taskId string) (*ViduResponse[*ViduTaskData], *types.OpenAIError) {
	// 先尝试新的官方查询接口
	queryResp, err := p.QueryCreations(taskId)
	if err != nil {
		return nil, err
	}

	// 将新格式转换为旧格式以保持兼容性
	response := &ViduResponse[*ViduTaskData]{
		TaskID:  taskId,
		Status:  queryResp.State,
		Message: queryResp.ErrCode,
		Credits: queryResp.Credits,
		Data: &ViduTaskData{
			TaskID:  taskId,
			Status:  queryResp.State,
			Credits: queryResp.Credits,
			Videos:  convertCreationsToVideos(queryResp.Creations),
		},
	}

	return response, nil
}

// 查询任务状态（新的官方接口）
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

// 辅助函数：将 Creations 转换为 ViduVideoResult 格式
func convertCreationsToVideos(creations []ViduCreation) []ViduVideoResult {
	videos := make([]ViduVideoResult, len(creations))
	for i, creation := range creations {
		videos[i] = ViduVideoResult{
			ID:  creation.ID,
			URL: creation.URL,
		}
	}
	return videos
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
    if err.Error.Code == "" {
        return nil
    }

	return &types.OpenAIError{
		Code:    err.Error.Code,
		Message: err.Error.Message,
		Type:    "vidu_error",
    }
}

// 将模型名规范化为官方API支持的格式。
// 根据官方文档，支持的模型格式为：viduq1, vidu1.5, vidu2.0, viduq1-classic
// 保持与官方API文档一致的模型名称格式。
func normalizeModelName(model string) string {
    m := strings.TrimSpace(strings.ToLower(model))
    switch m {
    case "vidu1.5":
        return "vidu1.5"  // 官方文档确认支持此格式
    case "vidu2.0":
        return "vidu2.0"  // 保持原格式
    case "viduq1":
        return "viduq1"   // 官方文档确认支持此格式
    case "viduq1-classic":
        return "viduq1-classic"  // 保持原格式
    // 支持连字符格式的向后兼容
    case "vidu-1.5":
        return "vidu1.5"
    case "vidu-2.0":
        return "vidu2.0"
    case "vidu-q1":
        return "viduq1"
    case "vidu-q1-classic":
        return "viduq1-classic"
    default:
        // 其他输入保持原样，避免阻断新模型接入
        return m
    }
}
