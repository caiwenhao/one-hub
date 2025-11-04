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
		SubmitPath: "/ent/v2/%s",                 // POST
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
	if p.isPolloUpstream() {
		// Pollo.ai 使用 x-api-key 认证
		if p.Channel.Key != "" {
			headers["x-api-key"] = p.Channel.Key
		}
	} else {
		if p.Channel.Key != "" {
			headers["Authorization"] = fmt.Sprintf("Token %s", p.Channel.Key)
		}
	}
	headers["Content-Type"] = "application/json"
	return headers
}

// 提交任务 - 对齐官方文档
func (p *ViduProvider) Submit(action string, request interface{}) (*ViduResponse, *types.OpenAIError) {
	// Pollo 上游：内部改写到 Pollo 路由与体
	if p.isPolloUpstream() {
		// 选择 endpoint 模型名（Pollo 对 Q1 使用 vidu-q1，其它沿用原名）
		modelName := ""
		switch v := request.(type) {
		case *ViduTaskRequest:
			modelName = normalizeModelName(v.Model)
		case *ViduReference2ImageRequest:
			modelName = normalizeModelName(v.Model)
		}
		endpointModel := toPolloModelEndpointName(modelName)
		polloURL := p.GetFullRequestURL(fmt.Sprintf("/generation/vidu/%s", endpointModel), "")
		headers := p.GetRequestHeaders()

		var reqBody any
		if v, ok := request.(*ViduTaskRequest); ok {
			// Pollo 能力约束：Q2 家族不支持纯文生（必须有 image）
			if action == ViduActionText2Video {
				low := strings.ToLower(endpointModel)
				if (strings.HasPrefix(low, "viduq2") || low == "viduq2" || strings.Contains(low, "q2")) && (v.Images == nil || len(v.Images) == 0) {
					return nil, &types.OpenAIError{Message: "Pollo upstream does not support text2video for Q2 models. Please use 'viduq1', or provide images (img2video/start-end2video), or switch upstream to official.", Type: "vidu_unsupported"}
				}
			}
			reqBody = buildPolloRequestFromVidu(action, v)
		} else if v2, ok := request.(*ViduReference2ImageRequest); ok {
			v := &ViduTaskRequest{
				Model:       v2.Model,
				Images:      v2.Images,
				Prompt:      v2.Prompt,
				Seed:        v2.Seed,
				AspectRatio: v2.AspectRatio,
				Payload:     v2.Payload,
				CallbackURL: v2.CallbackURL,
			}
			reqBody = buildPolloRequestFromVidu(action, v)
		} else {
			reqBody = request
		}

		req, err := p.Requester.NewRequest(http.MethodPost, polloURL, p.Requester.WithBody(reqBody), p.Requester.WithHeader(headers))
		if err != nil {
			return nil, &types.OpenAIError{Message: fmt.Sprintf("create request failed: %s", err.Error()), Type: "vidu_request_error"}
		}
		var polloResp struct{ TaskID, Status string }
		_, errWithCode := p.Requester.SendRequest(req, &polloResp, false)
		if errWithCode != nil {
			return nil, &errWithCode.OpenAIError
		}
		mapped := ViduResponse{TaskID: polloResp.TaskID, State: mapPolloStatus(polloResp.Status)}
		return &mapped, nil
	}

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
			// 先进行通用模型规范化
			model := normalizeModelName(taskReq.Model)
			// 再根据 action 做兼容修正：例如 text2video 的 Q2 家族需用通用 "viduq2"
			model = normalizeModelByAction(action, model)
			normalizedReq.Model = model
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
	if p.isPolloUpstream() {
		polloURL := p.GetFullRequestURL(fmt.Sprintf("/generation/%s/status", taskId), "")
		headers := p.GetRequestHeaders()
		req, err := p.Requester.NewRequest(http.MethodGet, polloURL, p.Requester.WithHeader(headers))
		if err != nil {
			return nil, &types.OpenAIError{Message: fmt.Sprintf("create request failed: %s", err.Error()), Type: "vidu_request_error"}
		}
		var polloResp struct{ TaskID, Status string }
		_, errWithCode := p.Requester.SendRequest(req, &polloResp, false)
		if errWithCode != nil {
			return nil, &errWithCode.OpenAIError
		}
		return &ViduQueryResponse{ID: polloResp.TaskID, State: mapPolloStatus(polloResp.Status)}, nil
	}
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
	if p.isPolloUpstream() {
		return nil, &types.OpenAIError{Message: "cancel is not supported by upstream: pollo.ai", Type: "vidu_request_error"}
	}
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
	// 简写/别名：无后缀的 q2 归一为 q2-pro
	case "viduq2", "vidu-q2", "q2":
		return ViduModelQ2Pro
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

// 是否选择了 Pollo 上游：根据 custom_parameter.upstream 或 base_url 判断
func (p *ViduProvider) isPolloUpstream() bool {
	// 优先 custom_parameter
	custom := p.Channel.GetCustomParameter()
	if strings.Contains(strings.ToLower(custom), "pollo") {
		return true
	}
	base := strings.ToLower(strings.TrimSpace(p.Channel.GetBaseURL()))
	if strings.Contains(base, "pollo.ai") {
		return true
	}
	return false
}

// 注：Pollo 上游在本集成中使用与 Vidu 官方一致的路径与请求体；仅鉴权头不同（x-api-key）。
// 根据 action 做模型的兼容修正
// 目前已知：text2video 在 Q2 家族下，上游更倾向使用通用 "viduq2"（而非 "viduq2-pro"/"viduq2-turbo"）
func normalizeModelByAction(action, model string) string {
	m := strings.TrimSpace(strings.ToLower(model))
	if action == ViduActionText2Video {
		if strings.HasPrefix(m, "viduq2-") || m == "viduq2" || m == "q2" || m == "vidu-q2" {
			return "viduq2"
		}
	}
	return model
}

// Pollo: 将 Vidu 请求映射为 Pollo input/webhook 结构
func buildPolloRequestFromVidu(action string, req *ViduTaskRequest) any {
	type polloInput struct {
		Image             string `json:"image,omitempty"`
		ImageTail         string `json:"imageTail,omitempty"`
		Prompt            string `json:"prompt,omitempty"`
		MovementAmplitude string `json:"movementAmplitude,omitempty"`
		Length            *int   `json:"length,omitempty"`
		Resolution        string `json:"resolution,omitempty"`
		Seed              *int   `json:"seed,omitempty"`
		GenerateAudio     *bool  `json:"generateAudio,omitempty"`
		Style             string `json:"style,omitempty"`
		AspectRatio       string `json:"aspectRatio,omitempty"`
	}
	body := struct {
		Input      polloInput `json:"input"`
		WebhookURL string     `json:"webhookUrl,omitempty"`
	}{}
	if len(req.Images) > 0 {
		body.Input.Image = req.Images[0]
		if len(req.Images) > 1 {
			body.Input.ImageTail = req.Images[1]
		}
	}
	body.Input.Prompt = req.Prompt
	if req.MovementAmplitude != "" {
		body.Input.MovementAmplitude = req.MovementAmplitude
	}
	body.Input.Length = req.Duration
	if req.Resolution != "" {
		body.Input.Resolution = req.Resolution
	}
	body.Input.Seed = req.Seed
	body.Input.GenerateAudio = req.BGM
	body.Input.Style = req.Style
	body.Input.AspectRatio = req.AspectRatio
	body.WebhookURL = req.CallbackURL
	return body
}

// Pollo 状态映射
func mapPolloStatus(status string) string {
	s := strings.ToLower(strings.TrimSpace(status))
	switch s {
	case "waiting":
		return ViduStatusQueueing
	case "processing":
		return ViduStatusProcessing
	case "succeed":
		return ViduStatusSuccess
	case "failed":
		return ViduStatusFailed
	default:
		return s
	}
}

// Pollo 模型 endpoint 命名差异：viduq1 -> vidu-q1
func toPolloModelEndpointName(model string) string {
	m := strings.TrimSpace(strings.ToLower(model))
	switch m {
	case "viduq1":
		return "vidu-q1"
	default:
		return m
	}
}

// GetModelList 仅返回文档中的基础模型列表，用于“创建渠道-模型自动填充/获取可用模型”
// 不返回任何精确计费变体，避免给运营带来维护成本。
func (p *ViduProvider) GetModelList() ([]string, error) {
	return []string{
		ViduModelQ2Pro,
		ViduModelQ2Turbo,
		ViduModelQ1,
		ViduModelQ1Classic,
		ViduModel20,
		ViduModel15,
	}, nil
}
