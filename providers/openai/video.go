package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"one-api/common"
	"one-api/common/config"
	"one-api/common/logger"
	"one-api/common/requester"
	"one-api/common/storage"
	"one-api/common/utils"
	"one-api/types"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	soraVendorOfficial = "official"
	soraVendorMountSea = "mountsea"
	soraVendorSutui    = "sutui"
	soraVendorApimart  = "apimart"
)

type openAIVideoResponse struct {
	types.VideoJob
	types.OpenAIErrorResponse
}

type openAIVideoListResponse struct {
	Object string           `json:"object"`
	Data   []types.VideoJob `json:"data"`
	types.OpenAIErrorResponse
}

type mountSeaCreateRequest struct {
	Model           string   `json:"model"`
	Prompt          string   `json:"prompt"`
	Duration        int      `json:"duration"`
	Size            string   `json:"size,omitempty"`
	Orientation     string   `json:"orientation,omitempty"`
	RemoveWatermark bool     `json:"removeWatermark,omitempty"`
	Images          []string `json:"images,omitempty"`
	InputReference  string   `json:"inputReference,omitempty"`
	RemixVideoID    string   `json:"remixVideoId,omitempty"`
	Seed            string   `json:"seed,omitempty"`
}

type mountSeaCreateResponse struct {
	TaskID      string `json:"taskId"`
	Status      string `json:"status"`
	ErrorCode   int    `json:"errorCode"`
	ErrorMsg    string `json:"errorMessage"`
	TraceID     string `json:"traceId"`
	CreatedAt   string `json:"createdAt"`
	Orientation string `json:"orientation"`
	Size        string `json:"size"`
}

type mountSeaVideoResult struct {
	VideoURL     string `json:"video_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type mountSeaTaskResult struct {
	TaskID      string               `json:"taskId"`
	Status      string               `json:"status"`
	Result      *mountSeaVideoResult `json:"result"`
	ErrorMsg    string               `json:"errorMessage"`
	ErrorCode   int                  `json:"errorCode"`
	FinishedAt  string               `json:"finishedAt"`
	CreatedAt   string               `json:"createdAt"`
	TraceID     string               `json:"traceId"`
	Orientation string               `json:"orientation"`
	Size        string               `json:"size"`
	Duration    int                  `json:"duration"`
}

type soraSizeInfo struct {
	Resolution  string
	Orientation string
	SizeLabel   string
}

type apimartCreateRequest struct {
	Model       string   `json:"model"`
	Prompt      string   `json:"prompt"`
	Duration    int      `json:"duration,omitempty"`
	AspectRatio string   `json:"aspect_ratio,omitempty"`
	ImageURLs   []string `json:"image_urls,omitempty"`
	Watermark   *bool    `json:"watermark,omitempty"`
}

type apimartCreateResponse struct {
	Code  int               `json:"code"`
	Data  []apimartTaskInfo `json:"data"`
	Error *apimartError     `json:"error"`
}

type apimartTaskInfo struct {
	Status string `json:"status"`
	TaskID string `json:"task_id"`
}

type apimartError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type apimartTaskResponse struct {
	Code  int              `json:"code"`
	Data  *apimartTaskData `json:"data"`
	Error *apimartError    `json:"error"`
}

type apimartTaskData struct {
	ID            string             `json:"id"`
	Status        string             `json:"status"`
	Progress      any                `json:"progress"`
	Result        *apimartTaskResult `json:"result"`
	Created       int64              `json:"created"`
	Completed     int64              `json:"completed"`
	EstimatedTime int64              `json:"estimated_time"`
	ActualTime    int64              `json:"actual_time"`
	Model         string             `json:"model"`
	Type          string             `json:"type"`
	TaskInfo      map[string]any     `json:"task_info"`
	Usage         map[string]any     `json:"usage"`
	Error         *apimartError      `json:"error"`
}

type apimartTaskResult struct {
	Videos []apimartTaskVideo `json:"videos"`
	Images []any              `json:"images"`
}

type apimartTaskVideo struct {
	URL        stringOrArray `json:"url"`
	Variant    string        `json:"variant"`
	Quality    string        `json:"quality"`
	Cover      string        `json:"cover"`
	ExpiresAt  int64         `json:"expires_at"`
	Resolution string        `json:"resolution"`
	Seconds    int           `json:"seconds"`
}

type stringOrArray string

func (s *stringOrArray) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*s = ""
		return nil
	}
	first := data[0]
	switch first {
	case '[':
		var arr []string
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		if len(arr) > 0 {
			*s = stringOrArray(arr[0])
		} else {
			*s = ""
		}
		return nil
	case 'n': // null
		*s = ""
		return nil
	default:
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		*s = stringOrArray(str)
		return nil
	}
}

func (s stringOrArray) String() string {
	return string(s)
}

func (p *OpenAIProvider) CreateVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	// NewAPI 渠道：将 OpenAI 标准 /v1/videos 请求映射为供应商 /v1/videos/generations
	if p.Channel != nil && p.Channel.Type == config.ChannelTypeNewAPI {
		return p.createNewAPIVideoFromOpenAI(request)
	}
	// 根据渠道配置与 BaseURL 识别供应商，分别适配官方与 MountSea 聚合
	switch p.detectSoraVendor() {
	case soraVendorMountSea:
		return p.createMountSeaVideo(request)
	case soraVendorSutui:
		return p.createSutuiVideo(request)
	case soraVendorApimart:
		return p.createApimartVideo(request)
	default:
		// 统一遵循 OpenAI 标准视频路由：/v1/videos
		return p.createOfficialVideo(request)
	}
}

func (p *OpenAIProvider) RetrieveVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	switch p.detectSoraVendor() {
	case soraVendorMountSea:
		return p.retrieveMountSeaVideo(videoID)
	case soraVendorSutui:
		return p.retrieveSutuiVideo(videoID)
	case soraVendorApimart:
		return p.retrieveApimartVideo(videoID)
	default:
		// 统一遵循 OpenAI 标准视频路由：/v1/videos
		return p.retrieveOfficialVideo(videoID)
	}
}

func (p *OpenAIProvider) DownloadVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	switch p.detectSoraVendor() {
	case soraVendorMountSea:
		return p.downloadMountSeaVideo(videoID, variant)
	case soraVendorSutui:
		return p.downloadSutuiVideo(videoID, variant)
	case soraVendorApimart:
		return p.downloadApimartVideo(videoID, variant)
	default:
		// 统一遵循 OpenAI 标准视频路由：/v1/videos
		return p.downloadOfficialVideo(videoID, variant)
	}
}

func (p *OpenAIProvider) detectSoraVendor() string {
	if p.Channel != nil && p.Channel.Plugin != nil {
		if plugin := p.Channel.Plugin.Data(); plugin != nil {
			if soraConfig, ok := plugin["sora"]; ok {
				if vendor, ok := soraConfig["vendor"].(string); ok && vendor != "" {
					return strings.ToLower(vendor)
				}
			}
		}
	}

	base := strings.ToLower(strings.TrimSpace(p.GetBaseURL()))
	if base != "" && strings.Contains(base, "mountsea") {
		return soraVendorMountSea
	}
	if base != "" && (strings.Contains(base, "sutui") || strings.Contains(base, "st-ai") || strings.Contains(base, "sora2.pub")) {
		return soraVendorSutui
	}
	if base != "" && strings.Contains(base, "apimart") {
		return soraVendorApimart
	}

	return soraVendorOfficial
}

func (p *OpenAIProvider) createOfficialVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	// 若为 multipart/form-data，则透传原始请求体，保持与官方一致
	contentType := ""
	if p.Context != nil && p.Context.Request != nil {
		contentType = p.Context.Request.Header.Get("Content-Type")
	}

    if strings.Contains(strings.ToLower(contentType), "multipart/form-data") {
        urlPath, errWithCode := p.GetSupportedAPIUri(config.RelayModeOpenAIVideo)
        if errWithCode != nil {
            return nil, errWithCode
        }
        fullURL := p.GetFullRequestURL(urlPath, request.Model)
        // 透传原始 body
        raw, ok := p.GetRawBody()
        if !ok {
            return nil, common.StringErrorWrapperLocal("missing raw multipart body", "missing_body", http.StatusBadRequest)
        }
        headers := p.GetRequestHeaders()
        // 关键修复：将下游原始 Content-Type（含 boundary）透传到上游
        if ct := strings.TrimSpace(contentType); ct != "" {
            headers["Content-Type"] = ct
        }

		// 调试输出：解析 multipart 中的 seconds 字段，辅助定位上游严格校验失败
        func() {
            // 使用已透传的 contentType 解析 seconds，避免 headers 中缺失导致日志/校验失真
            ct := headers["Content-Type"]
            secondsVal := ""
            if ct != "" {
                if mt, params, e := mime.ParseMediaType(ct); e == nil && strings.HasPrefix(strings.ToLower(mt), "multipart/") {
                    if boundary := params["boundary"]; strings.TrimSpace(boundary) != "" {
                        mr := multipart.NewReader(bytes.NewReader(raw), boundary)
                        for {
                            part, e2 := mr.NextPart()
                            if e2 == io.EOF { break }
                            if e2 != nil { break }
                            name := part.FormName()
                            if name == "seconds" && part.FileName() == "" {
                                b, _ := io.ReadAll(part)
                                secondsVal = strings.TrimSpace(string(b))
                                _ = part.Close()
                                break
                            }
                            _ = part.Close()
                        }
                    }
                }
            }
            logger.SysDebug(fmt.Sprintf("video.create.multipart debug -> ct=%s seconds_field=%q raw_len=%d", ct, secondsVal, len(raw)))
        }()
        // EzlinkAI: 严格秒数校验（仅允许 4/8/12），不做宽松归一化
        var httpReq *http.Request
        base := strings.ToLower(strings.TrimSpace(p.GetBaseURL()))
        if strings.Contains(base, "ezlinkai.com") {
            if ct := headers["Content-Type"]; strings.TrimSpace(ct) != "" {
                if err := p.validateMultipartSecondsEzlinkAI(raw, ct); err != nil {
                    logger.SysDebug(fmt.Sprintf("video.create.multipart validate_seconds failed -> err=%s", err.Error()))
                    return nil, common.StringErrorWrapperLocal(err.Error(), "invalid_seconds", http.StatusBadRequest)
                }
            }
        }
        if req2, e := p.Requester.NewRequest(
            http.MethodPost,
            fullURL,
            p.Requester.WithBody(raw),
            p.Requester.WithHeader(headers),
        ); e == nil {
            httpReq = req2
        } else {
            return nil, common.ErrorWrapper(e, "new_request_failed", http.StatusInternalServerError)
        }
        // 最佳实践：同步下游 Content-Length，必要时回落为原始字节长度
        if p.Context != nil && p.Context.Request != nil && p.Context.Request.ContentLength > 0 {
            httpReq.ContentLength = p.Context.Request.ContentLength
        } else {
            httpReq.ContentLength = int64(len(raw))
        }
        defer httpReq.Body.Close()

		response := &openAIVideoResponse{}
		_, errWithCode = p.Requester.SendRequest(httpReq, response, false)
		if errWithCode != nil {
			return nil, errWithCode
		}
		if openErr := ErrorHandle(&response.OpenAIErrorResponse); openErr != nil {
			return nil, &types.OpenAIErrorWithStatusCode{OpenAIError: *openErr, StatusCode: http.StatusBadRequest}
		}
		job := response.VideoJob
		if job.Object == "" {
			job.Object = "video"
		}
		if job.Model == "" {
			job.Model = request.Model
		}
		if job.CreatedAt == 0 {
			job.CreatedAt = time.Now().Unix()
		}
		if job.Progress == 0 {
			// 与官方示例保持一致
			job.Progress = 0
		}
		if job.Quality == "" {
			job.Quality = "standard"
		}
		return &job, nil
	}

	// 默认 JSON 方式
	reqCopy := *request
	// EzlinkAI 上游仅支持 seconds: 4 / 8 / 12，严格校验
	base := strings.ToLower(strings.TrimSpace(p.GetBaseURL()))
	if strings.Contains(base, "ezlinkai.com") {
		switch reqCopy.Seconds {
		case 4, 8, 12:
			// ok
		default:
			return nil, common.StringErrorWrapperLocal("seconds must be one of 4, 8, 12 for ezlinkai", "invalid_seconds", http.StatusBadRequest)
		}
	}
	httpReq, errWithCode := p.GetRequestTextBody(config.RelayModeOpenAIVideo, reqCopy.Model, &reqCopy)
	if errWithCode != nil {
		return nil, errWithCode
	}
	defer httpReq.Body.Close()

	response := &openAIVideoResponse{}
	_, errWithCode = p.Requester.SendRequest(httpReq, response, false)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if openErr := ErrorHandle(&response.OpenAIErrorResponse); openErr != nil {
		return nil, &types.OpenAIErrorWithStatusCode{
			OpenAIError: *openErr,
			StatusCode:  http.StatusBadRequest,
		}
	}

	job := response.VideoJob
	if job.Object == "" {
		job.Object = "video"
	}
	if job.Model == "" {
		job.Model = request.Model
	}
	if job.CreatedAt == 0 {
		job.CreatedAt = time.Now().Unix()
	}
	if job.Progress == 0 {
		job.Progress = 0
	}
	if job.Quality == "" {
		job.Quality = "standard"
	}
	return &job, nil
}

// --- NewAPI adapter: map OpenAI /v1/videos -> vendor /v1/videos/generations ---
type newAPIVideoGenerationsReq struct {
	Model       string   `json:"model"`
	Prompt      string   `json:"prompt,omitempty"`
	Duration    int      `json:"duration,omitempty"`
	AspectRatio string   `json:"aspect_ratio,omitempty"`
	ImageURLs   []string `json:"image_urls,omitempty"`
	Watermark   *bool    `json:"watermark,omitempty"`
}

type newAPICreateResp struct {
	Code int `json:"code"`
	Data []struct {
		Status string `json:"status"`
		TaskID string `json:"task_id"`
	} `json:"data"`
	Error any `json:"error"`
}

func (p *OpenAIProvider) createNewAPIVideoFromOpenAI(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	// 构造 generations 请求体
	body := &newAPIVideoGenerationsReq{Model: request.Model, Prompt: request.Prompt}
	if request.Seconds > 0 {
		body.Duration = request.Seconds
	}
	// 映射 size -> aspect_ratio
	if s := strings.TrimSpace(request.Size); s != "" {
		// 简单按常用分辨率判断
		switch strings.ToLower(s) {
		case "1280x720":
			body.AspectRatio = "16:9"
		case "720x1280":
			body.AspectRatio = "9:16"
		case "1792x1024":
			body.AspectRatio = "16:9"
		case "1024x1792":
			body.AspectRatio = "9:16"
		}
	}
	// 图生视频：聚合输入图
	if len(request.InputImages) > 0 {
		body.ImageURLs = append(body.ImageURLs, request.InputImages...)
	}
	if strings.TrimSpace(request.InputImage) != "" {
		body.ImageURLs = append(body.ImageURLs, request.InputImage)
	}
	// 水印：与 remove_watermark 反向
	if request.RemoveWatermark {
		f := false
		body.Watermark = &f
	}

	headers := p.GetRequestHeaders()
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"

	fullURL := strings.TrimSuffix(p.GetBaseURL(), "/") + "/v1/videos/generations"
	req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

	// 对于第三方 JSON，关闭 OpenAI 错误包装前缀
	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	resp := &newAPICreateResp{}
	if _, errWith := p.Requester.SendRequest(req, resp, false); errWith != nil {
		return nil, errWith
	}
	// 解析返回
	var taskID, status string
	if len(resp.Data) > 0 {
		taskID = resp.Data[0].TaskID
		status = resp.Data[0].Status
	}
	if taskID == "" {
		return nil, common.StringErrorWrapperLocal("invalid upstream response", "bad_upstream", http.StatusBadGateway)
	}
	// 映射 status：submitted -> queued
	st := strings.ToLower(strings.TrimSpace(status))
	if st == "submitted" {
		st = "queued"
	}

	job := &types.VideoJob{
		ID:        taskID,
		Object:    "video",
		Model:     request.Model,
		Status:    st,
		Progress:  0,
		Seconds:   request.Seconds,
		Size:      request.Size,
		Quality:   "standard",
		CreatedAt: time.Now().Unix(),
	}
	return job, nil
}

func (p *OpenAIProvider) retrieveOfficialVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	fullURL := p.buildOfficialVideoURL(videoID, "")
	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(p.GetRequestHeaders()))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}

	response := &openAIVideoResponse{}
	_, errWithCode := p.Requester.SendRequest(req, response, false)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if openErr := ErrorHandle(&response.OpenAIErrorResponse); openErr != nil {
		return nil, &types.OpenAIErrorWithStatusCode{
			OpenAIError: *openErr,
			StatusCode:  http.StatusBadRequest,
		}
	}

	job := response.VideoJob
	if job.Object == "" {
		job.Object = "video"
	}
	if job.Quality == "" {
		job.Quality = "standard"
	}
	return &job, nil
}

func (p *OpenAIProvider) downloadOfficialVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	path := fmt.Sprintf("%s/%s/content", strings.TrimSuffix(p.Config.Videos, "/"), videoID)
	fullURL := p.buildOfficialVideoURL("", path)

	if variant != "" && variant != "video" {
		query := url.Values{}
		query.Set("variant", variant)
		if strings.Contains(fullURL, "?") {
			fullURL = fullURL + "&" + query.Encode()
		} else {
			fullURL = fullURL + "?" + query.Encode()
		}
	}

	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(p.GetRequestHeaders()))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}

	resp, errWithCode := p.Requester.SendRequest(req, nil, true)
	if errWithCode != nil {
		return nil, errWithCode
	}

	return resp, nil
}

// RemixVideo: 官方支持 /v1/videos/{id}/remix；MountSea 回退为通过 create 接口的 remixVideoId 能力
func (p *OpenAIProvider) RemixVideo(videoID string, prompt string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	switch p.detectSoraVendor() {
	case soraVendorMountSea:
		// 通过 create + RemixVideoID 实现
		req := &types.VideoCreateRequest{RemixVideoID: videoID, Prompt: prompt, Model: p.GetOriginalModel()}
		return p.createMountSeaVideo(req)
	case soraVendorSutui:
		// 通过 create + RemixVideoID 实现
		req := &types.VideoCreateRequest{RemixVideoID: videoID, Prompt: prompt, Model: p.GetOriginalModel()}
		return p.createSutuiVideo(req)
	case soraVendorApimart:
		return p.remixApimartVideo(videoID, prompt)
	default:
		// 官方路径
		basePath := strings.TrimSuffix(p.Config.Videos, "/")
		path := fmt.Sprintf("%s/%s/remix", basePath, videoID)
		fullURL := p.buildOfficialVideoURL("", path)
		headers := p.GetRequestHeaders()
		body := map[string]string{"prompt": prompt}
		req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
		if err != nil {
			return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
		}
		if req.Body != nil {
			defer req.Body.Close()
		}
		response := &openAIVideoResponse{}
		_, errWithCode := p.Requester.SendRequest(req, response, false)
		if errWithCode != nil {
			return nil, errWithCode
		}
		if openErr := ErrorHandle(&response.OpenAIErrorResponse); openErr != nil {
			return nil, &types.OpenAIErrorWithStatusCode{OpenAIError: *openErr, StatusCode: http.StatusBadRequest}
		}
		job := response.VideoJob
		if job.Object == "" {
			job.Object = "video"
		}
		return &job, nil
	}
}

// ListVideos: 官方 /v1/videos?after=&limit=&order=
func (p *OpenAIProvider) ListVideos(after string, limit int, order string) (*types.VideoList, *types.OpenAIErrorWithStatusCode) {
	if p.detectSoraVendor() != soraVendorOfficial {
		return nil, common.StringErrorWrapperLocal("list videos is not supported for this channel", "unsupported_api", http.StatusNotImplemented)
	}
	base := strings.TrimSuffix(p.GetFullRequestURL("", ""), "/")
	videosPath := strings.TrimSuffix(p.Config.Videos, "/")
	fullURL := fmt.Sprintf("%s%s", base, videosPath)
	q := url.Values{}
	if strings.TrimSpace(after) != "" {
		q.Set("after", after)
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if strings.TrimSpace(order) != "" {
		q.Set("order", order)
	}
	if enc := q.Encode(); enc != "" {
		fullURL = fullURL + "?" + enc
	}
	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(p.GetRequestHeaders()))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}
	response := &openAIVideoListResponse{}
	_, errWithCode := p.Requester.SendRequest(req, response, false)
	if errWithCode != nil {
		return nil, errWithCode
	}
	if openErr := ErrorHandle(&response.OpenAIErrorResponse); openErr != nil {
		return nil, &types.OpenAIErrorWithStatusCode{OpenAIError: *openErr, StatusCode: http.StatusBadRequest}
	}
	return &types.VideoList{Object: response.Object, Data: response.Data}, nil
}

// DeleteVideo: 官方 DELETE /v1/videos/{id}
func (p *OpenAIProvider) DeleteVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	if p.detectSoraVendor() != soraVendorOfficial {
		return nil, common.StringErrorWrapperLocal("delete video is not supported for this channel", "unsupported_api", http.StatusNotImplemented)
	}
	basePath := strings.TrimSuffix(p.Config.Videos, "/")
	fullURL := p.buildOfficialVideoURL("", fmt.Sprintf("%s/%s", basePath, videoID))
	req, err := p.Requester.NewRequest(http.MethodDelete, fullURL, p.Requester.WithHeader(p.GetRequestHeaders()))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}
	response := &openAIVideoResponse{}
	_, errWithCode := p.Requester.SendRequest(req, response, false)
	if errWithCode != nil {
		return nil, errWithCode
	}
	if openErr := ErrorHandle(&response.OpenAIErrorResponse); openErr != nil {
		return nil, &types.OpenAIErrorWithStatusCode{OpenAIError: *openErr, StatusCode: http.StatusBadRequest}
	}
	job := response.VideoJob
	if job.Object == "" {
		job.Object = "video"
	}
	return &job, nil
}

func (p *OpenAIProvider) buildOfficialVideoURL(videoID string, overridePath string) string {
	baseURL := strings.TrimSuffix(p.GetFullRequestURL("", ""), "/")
	if overridePath != "" {
		overridePath = strings.TrimPrefix(overridePath, "/")
		return fmt.Sprintf("%s/%s", baseURL, overridePath)
	}

	videosPath := strings.TrimSuffix(p.Config.Videos, "/")
	if videosPath == "" {
		videosPath = "/v1/videos"
	}
	videoID = strings.TrimPrefix(videoID, "/")
	return fmt.Sprintf("%s%s/%s", baseURL, videosPath, videoID)
}

func (p *OpenAIProvider) createMountSeaVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	sizeInfo := normalizeSoraSize(request.Size)
	duration := request.Seconds
	if duration <= 0 {
		duration = 4
	}

	msReq := &mountSeaCreateRequest{
		Model:           request.Model,
		Prompt:          request.Prompt,
		Duration:        duration,
		Size:            sizeInfo.SizeLabel,
		Orientation:     sizeInfo.Orientation,
		RemoveWatermark: request.RemoveWatermark,
	}

	if len(request.InputImages) > 0 {
		msReq.Images = append(msReq.Images, request.InputImages...)
	}

	if request.InputImage != "" {
		msReq.Images = append(msReq.Images, request.InputImage)
	}

	if request.InputReference != "" {
		msReq.InputReference = request.InputReference
	}

	if request.RemixVideoID != "" {
		msReq.RemixVideoID = request.RemixVideoID
	}

	if request.Seed != "" {
		msReq.Seed = request.Seed
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", p.Channel.Key),
	}

	fullURL := strings.TrimSuffix(p.GetBaseURL(), "/") + "/sora/video/generate"
	req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(msReq), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() {
		p.Requester.IsOpenAI = originalOpenAIFlag
	}()

	response := &mountSeaCreateResponse{}
	_, errWithCode := p.Requester.SendRequest(req, response, false)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if response.TaskID == "" {
		return nil, common.ErrorWrapper(fmt.Errorf("missing taskId in mountsea response"), "mountsea_error", http.StatusBadGateway)
	}

	status, progress := mapMountSeaStatus(response.Status)
	job := &types.VideoJob{
		ID:       response.TaskID,
		Object:   "video",
		Model:    request.Model,
		Status:   status,
		Progress: progress,
		Seconds:  duration,
		Size:     sizeInfo.Resolution,
	}

	if response.CreatedAt != "" {
		if ts, err := parseMountSeaTime(response.CreatedAt); err == nil {
			job.CreatedAt = ts
		}
	}

	// 对齐官方：若缺失，填充 created_at、quality
	if job.CreatedAt == 0 {
		job.CreatedAt = time.Now().Unix()
	}
	if job.Quality == "" {
		job.Quality = "standard"
	}

	return job, nil
}

func (p *OpenAIProvider) retrieveMountSeaVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", p.Channel.Key),
	}

	fullURL := strings.TrimSuffix(p.GetBaseURL(), "/") + "/sora/task/result?taskId=" + url.QueryEscape(videoID)
	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() {
		p.Requester.IsOpenAI = originalOpenAIFlag
	}()

	response := &mountSeaTaskResult{}
	_, errWithCode := p.Requester.SendRequest(req, response, false)
	if errWithCode != nil {
		return nil, errWithCode
	}

	status, progress := mapMountSeaStatus(response.Status)
	job := &types.VideoJob{
		ID:       videoID,
		Object:   "video",
		Status:   status,
		Progress: progress,
	}

	if response.Result != nil {
		job.Result = &types.VideoJobResult{
			VideoURL:     response.Result.VideoURL,
			ThumbnailURL: response.Result.ThumbnailURL,
		}
	}

	if response.CreatedAt != "" {
		if ts, err := parseMountSeaTime(response.CreatedAt); err == nil {
			job.CreatedAt = ts
		}
	}

	if response.Duration > 0 {
		job.Seconds = response.Duration
	}

	if response.Size != "" && response.Orientation != "" {
		job.Size = formatResolutionByOrientation(response.Size, response.Orientation)
	} else if response.Size != "" {
		job.Size = response.Size
	}

	if response.ErrorCode != 0 || response.ErrorMsg != "" {
		job.Error = &types.VideoJobError{
			Code:    response.ErrorCode,
			Message: response.ErrorMsg,
		}
	}

	return job, nil
}

func (p *OpenAIProvider) downloadMountSeaVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	if variant != "" && variant != "video" {
		return nil, common.StringErrorWrapperLocal("variant not supported for MountSea channel", "unsupported_variant", http.StatusNotImplemented)
	}

	job, errWithCode := p.retrieveMountSeaVideo(videoID)
	if errWithCode != nil {
		return nil, errWithCode
	}

	if job.Result == nil || job.Result.VideoURL == "" {
		return nil, common.StringErrorWrapperLocal("video not ready", "video_not_ready", http.StatusAccepted)
	}

	req, err := p.Requester.NewRequest(http.MethodGet, job.Result.VideoURL, p.Requester.WithHeader(map[string]string{}))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() {
		p.Requester.IsOpenAI = originalOpenAIFlag
	}()

	resp, errWithCode := p.Requester.SendRequest(req, nil, true)
	if errWithCode != nil {
		return nil, errWithCode
	}

	return resp, nil
}

func (p *OpenAIProvider) createApimartVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
    // 与官方对齐：不对外部请求做刚性秒数限制，内部归一化为上游可接受档位
    modelName := strings.ToLower(strings.TrimSpace(request.Model))
    secs := normalizeApimartSeconds(modelName, request.Seconds)

    // 收集参考图 URL：支持 JSON（input_image(s)）与 multipart 文件（input_reference/input_images）
    collectedURLs := []string{}
    // JSON 直传 URL
    if len(request.InputImages) > 0 {
        collectedURLs = append(collectedURLs, request.InputImages...)
    }
    if strings.TrimSpace(request.InputImage) != "" {
        collectedURLs = append(collectedURLs, request.InputImage)
    }

    // 尝试解析 multipart（若为 multipart 则把文件上传到 S3 并生成 URL）
    contentType := ""
    if p.Context != nil && p.Context.Request != nil {
        contentType = p.Context.Request.Header.Get("Content-Type")
    }
    var (
        overrideModel   = ""
        overridePrompt  = ""
        overrideSeconds = 0
        overrideSize    = ""
        removeWatermark = request.RemoveWatermark
    )
    if strings.Contains(strings.ToLower(contentType), "multipart/form-data") {
        raw, ok := p.GetRawBody()
        if ok {
            mediaType, params, err := mime.ParseMediaType(contentType)
            if err == nil && strings.HasPrefix(strings.ToLower(mediaType), "multipart/") {
                if boundary := strings.TrimSpace(params["boundary"]); boundary != "" {
                    mr := multipart.NewReader(bytes.NewReader(raw), boundary)
                    const maxImageSize = 10 * 1024 * 1024
                    allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
                    for {
                        part, e := mr.NextPart()
                        if e == io.EOF { break }
                        if e != nil { break }
                        name := part.FormName()
                        if name == "" { _ = part.Close(); continue }
                        if fn := part.FileName(); fn != "" {
                            // 仅处理 input_reference / input_images 文件
                            low := strings.ToLower(name)
                            if low == "input_reference" || low == "input_images" {
                                b, e2 := io.ReadAll(part)
                                _ = part.Close()
                                if e2 != nil { continue }
                                if len(b) == 0 || len(b) > maxImageSize { continue }
                                ext := strings.ToLower(filepath.Ext(fn))
                                if !allowedExt[ext] { ext = ".jpg" }
                                key := utils.GetUUID() + ext
                                if url := storage.Upload(b, key); strings.TrimSpace(url) != "" {
                                    collectedURLs = append(collectedURLs, url)
                                }
                            } else {
                                _ = part.Close()
                            }
                            continue
                        }
                        // 文本字段
                        bv, _ := io.ReadAll(part)
                        _ = part.Close()
                        val := strings.TrimSpace(string(bv))
                        switch strings.ToLower(name) {
                        case "model":
                            overrideModel = val
                        case "prompt":
                            overridePrompt = val
                        case "seconds":
                            if v, err := strconv.Atoi(val); err == nil { overrideSeconds = v }
                        case "size":
                            overrideSize = val
                        case "input_image", "input_images":
                            if val != "" { collectedURLs = append(collectedURLs, val) }
                        case "remove_watermark":
                            if strings.EqualFold(val, "true") || val == "1" { removeWatermark = true }
                        }
                    }
                }
            }
        }
    }

    // 应用覆盖字段
    modelForUp := request.Model
    if strings.TrimSpace(overrideModel) != "" { modelForUp = overrideModel }
    promptForUp := request.Prompt
    if strings.TrimSpace(overridePrompt) != "" { promptForUp = overridePrompt }
    secondsForUp := secs
    if overrideSeconds > 0 { secondsForUp = overrideSeconds }
    // 再次归一化（兼容 multipart 覆盖）
    secondsForUp = normalizeApimartSeconds(modelForUp, secondsForUp)
    sizeForUp := request.Size
    if strings.TrimSpace(overrideSize) != "" { sizeForUp = overrideSize }

    body := &apimartCreateRequest{
        Model:  modelForUp,
        Prompt: promptForUp,
        Duration: secondsForUp,
    }
    if ar := buildApimartAspectRatio(sizeForUp); ar != "" { body.AspectRatio = ar }
    if len(collectedURLs) > 0 { body.ImageURLs = append(body.ImageURLs, collectedURLs...) }
    if removeWatermark { f := false; body.Watermark = &f }

    // Debug：打印 apimart generations 关键入参（不打印具体 URL 以免泄漏）
    logger.SysDebug(fmt.Sprintf("apimart.gen debug -> model=%s duration=%d ar=%s image_urls=%d", body.Model, body.Duration, body.AspectRatio, len(body.ImageURLs)))

	headers := p.GetRequestHeaders()
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"

	fullURL := strings.TrimSuffix(p.GetBaseURL(), "/") + "/v1/videos/generations"
	req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() {
		p.Requester.IsOpenAI = originalOpenAIFlag
	}()

	resp := &apimartCreateResponse{}
	if _, errWith := p.Requester.SendRequest(req, resp, false); errWith != nil {
		return nil, errWith
	}
	if resp.Code != http.StatusOK {
		return nil, apimartErrorToOpenAI(resp.Error, resp.Code)
	}
	if len(resp.Data) == 0 || strings.TrimSpace(resp.Data[0].TaskID) == "" {
		return nil, common.StringErrorWrapperLocal("invalid upstream response", "bad_upstream", http.StatusBadGateway)
	}

    task := resp.Data[0]
    job := &types.VideoJob{
        ID:        task.TaskID,
        Object:    "video",
        Model:     request.Model,
        Status:    mapApimartStatus(task.Status),
        Progress:  0,
        Seconds:   secondsForUp,
        Size:      request.Size,
        Quality:   "standard",
        CreatedAt: time.Now().Unix(),
        Prompt:    request.Prompt,
    }
	if job.Status == "" {
		job.Status = "queued"
	}
	if job.Size == "" {
		if info := normalizeSoraSize(request.Size); info.Resolution != "" {
			job.Size = info.Resolution
		}
	}
	return job, nil
}

// --- Sutui (速推) upstream adapter ---

type sutuiCreateResponse struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Model     string `json:"model"`
	Status    string `json:"status"`
	Progress  any    `json:"progress"`
	CreatedAt int64  `json:"created_at"`
	Seconds   int    `json:"seconds"`
	Size      string `json:"size"`
}

type sutuiRetrieveResponse struct {
	CompletedAt        int64  `json:"completed_at"`
	CreatedAt          int64  `json:"created_at"`
	ID                 string `json:"id"`
	Model              string `json:"model"`
	Object             string `json:"object"`
	Progress           any    `json:"progress"`
	Seconds            int    `json:"seconds"`
	Size               string `json:"size"`
	Status             string `json:"status"`
	VideoURL           string `json:"video_url"`
	RemixedFromVideoID any    `json:"remixed_from_video_id"`
}

func anyToFloat(v any) float64 {
	switch t := v.(type) {
	case nil:
		return 0
	case float64:
		return t
	case float32:
		return float64(t)
	case int:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case string:
		if f, err := strconv.ParseFloat(t, 64); err == nil {
			return f
		}
		return 0
	default:
		return 0
	}
}

func (p *OpenAIProvider) createSutuiVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	// 严格校验秒数
	secs := request.Seconds
	mdl := strings.ToLower(strings.TrimSpace(p.GetOriginalModel()))
	if mdl == "" {
		mdl = strings.ToLower(strings.TrimSpace(request.Model))
	}
	if mdl == "sora-2-pro" {
		if !(secs == 15 || secs == 25) {
			return nil, common.StringErrorWrapperLocal("seconds must be 15 or 25 for sutui (sora-2-pro)", "invalid_seconds", http.StatusBadRequest)
		}
	} else {
		if !(secs == 10 || secs == 15) {
			return nil, common.StringErrorWrapperLocal("seconds must be 10 or 15 for sutui (sora-2)", "invalid_seconds", http.StatusBadRequest)
		}
	}

	// 构造 URL
	urlPath, errWithCode := p.GetSupportedAPIUri(config.RelayModeOpenAIVideo)
	if errWithCode != nil {
		return nil, errWithCode
	}
	fullURL := p.GetFullRequestURL(urlPath, request.Model)

	// 透传原始 Content-Type，尤其是 multipart 边界
	headers := p.GetRequestHeaders()
	contentType := ""
	if p.Context != nil && p.Context.Request != nil {
		contentType = p.Context.Request.Header.Get("Content-Type")
	}
	if contentType != "" {
		headers["Content-Type"] = contentType
	}

	// 临时关闭 OpenAI 错误前缀
	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	var httpReq *http.Request
	// 根据“宽松秒数策略”与 size 推断供应商 SKU（仅在 JSON 模式可安全改写入参）
	originalModel := strings.ToLower(strings.TrimSpace(p.GetOriginalModel()))
	if strings.Contains(strings.ToLower(contentType), "multipart/form-data") {
		raw, ok := p.GetRawBody()
		if !ok {
			return nil, common.StringErrorWrapperLocal("missing raw multipart body", "missing_body", http.StatusBadRequest)
		}
		// 对 multipart 进行改写：应用宽松秒数策略，动态选择 sutui SKU，并保持文件字段透传
		bodyReader, newCT, err := p.rewriteMultipartForSutui(raw, headers["Content-Type"], originalModel)
		if err == nil && bodyReader != nil && strings.TrimSpace(newCT) != "" {
			headers["Content-Type"] = newCT
			req, e2 := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(bodyReader), p.Requester.WithHeader(headers))
			if e2 != nil {
				return nil, common.ErrorWrapper(e2, "new_request_failed", http.StatusInternalServerError)
			}
			httpReq = req
		} else {
			// 回退：原样透传
			req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(raw), p.Requester.WithHeader(headers))
			if err != nil {
				return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
			}
			httpReq = req
		}
	} else {
		// JSON 输入：为保证与 Sutui 上游兼容，改造为 multipart/form-data 后转发
		reqCopy := *request
		if originalModel == "sora-2" || originalModel == "sora-2-pro" {
			sku, slotSec := pickSutuiSKU(originalModel, strings.TrimSpace(reqCopy.Size), reqCopy.Seconds)
			if sku != "" {
				reqCopy.Model = sku
				if slotSec > 0 {
					reqCopy.Seconds = slotSec
				}
			}
		}
		// 构建 multipart（仅文本字段）
		var buf bytes.Buffer
		builder := p.Requester.CreateFormBuilder(&buf)
		// 必填
		_ = builder.WriteField("model", strings.TrimSpace(reqCopy.Model))
		if strings.TrimSpace(reqCopy.Prompt) != "" {
			_ = builder.WriteField("prompt", reqCopy.Prompt)
		}
		if reqCopy.Seconds > 0 {
			_ = builder.WriteField("seconds", strconv.Itoa(reqCopy.Seconds))
		}
		if strings.TrimSpace(reqCopy.Size) != "" {
			_ = builder.WriteField("size", reqCopy.Size)
		}
		// 参考图 URL：下载后作为文件字段追加（默认不镜像，直接直传）
		const sutuiMaxImage = 10 * 1024 * 1024
		urlFilesAppended := 0
		appendURL := func(u string) {
			u = strings.TrimSpace(u)
			if u == "" { return }
			if data, filename, e := fetchImageForSutui(u, sutuiMaxImage); e == nil && len(data) > 0 {
				if err := builder.CreateFormFileReader("input_images", bytes.NewReader(data), filename); err == nil {
					urlFilesAppended++
				}
			}
		}
		if len(request.InputImages) > 0 {
			for _, u := range request.InputImages { appendURL(u) }
		}
		if strings.TrimSpace(request.InputImage) != "" { appendURL(request.InputImage) }

		// 其它可选参数（若将来 Sutui 支持可扩展）
		_ = builder.Close()
		// Debug: 仅记录成功追加的 URL 文件数量
		logger.SysDebug(fmt.Sprintf("sutui.gen debug -> url_files_appended=%d", urlFilesAppended))
		headers["Content-Type"] = builder.FormDataContentType()
		req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(&buf), p.Requester.WithHeader(headers))
		if err != nil {
			return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
		}
		httpReq = req
	}
	if httpReq.Body != nil {
		defer httpReq.Body.Close()
	}

	respObj := &sutuiCreateResponse{}
	_, errWith := p.Requester.SendRequest(httpReq, respObj, false)
	if errWith != nil {
		return nil, errWith
	}

	job := &types.VideoJob{
		ID:        respObj.ID,
		Object:    respObj.Object,
		CreatedAt: respObj.CreatedAt,
		Status:    respObj.Status,
		// 对外只暴露官方模型名
		Model:    p.GetOriginalModel(),
		Progress: anyToFloat(respObj.Progress),
		Seconds:  respObj.Seconds,
		Size:     respObj.Size,
		Quality:  "standard",
	}
	if job.Object == "" {
		job.Object = "video"
	}
	if job.CreatedAt == 0 {
		job.CreatedAt = time.Now().Unix()
	}
	// 若上游返回 720x720（Sutui SD 默认尺寸），根据请求 size 推断并覆盖为官方分辨率
	if strings.TrimSpace(job.Size) == "" || strings.TrimSpace(job.Size) == "720x720" {
		info := normalizeSoraSize(request.Size)
		if strings.TrimSpace(info.Resolution) != "" {
			job.Size = info.Resolution
		}
	}
	return job, nil
}

// rewriteMultipartForSutui 解析并重建 multipart/form-data，请求体中动态改写 model/seconds，保持其余字段与文件不变
func (p *OpenAIProvider) rewriteMultipartForSutui(raw []byte, contentType string, originalModel string) (io.Reader, string, error) {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil || !strings.HasPrefix(strings.ToLower(mediaType), "multipart/") {
		return nil, "", fmt.Errorf("invalid multipart content-type")
	}
	boundary := params["boundary"]
	if strings.TrimSpace(boundary) == "" {
		return nil, "", fmt.Errorf("missing boundary")
	}

	// 解析原始 multipart
	mr := multipart.NewReader(bytes.NewReader(raw), boundary)
	// 暂存字段与文件
	type filePart struct {
		field    string
		filename string
		data     []byte
	}
	fields := map[string][]string{}
	files := []filePart{}
	urlFilesAppended := 0
	for {
		part, e := mr.NextPart()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, "", e
		}
		name := part.FormName()
		if name == "" {
			// 跳过无名字段
			_ = part.Close()
			continue
		}
		if fn := part.FileName(); fn != "" {
			b, e2 := io.ReadAll(part)
			_ = part.Close()
			if e2 != nil {
				return nil, "", e2
			}
			files = append(files, filePart{field: name, filename: fn, data: b})
			continue
		}
		b, e3 := io.ReadAll(part)
		_ = part.Close()
		if e3 != nil {
			return nil, "", e3
		}
		fields[name] = append(fields[name], string(b))
	}

	// 读取外部字段
	size := ""
	if arr, ok := fields["size"]; ok && len(arr) > 0 {
		size = strings.TrimSpace(arr[0])
	}
	seconds := 0
	if arr, ok := fields["seconds"]; ok && len(arr) > 0 {
		if v, err := strconv.Atoi(strings.TrimSpace(arr[0])); err == nil {
			seconds = v
		}
	}

	// 将文本字段中的 URL 参考图下载为文件并追加
	const maxImageSize = 10 * 1024 * 1024
	appendURL := func(u string) {
		u = strings.TrimSpace(u)
		if u == "" { return }
		if data, filename, e := fetchImageForSutui(u, maxImageSize); e == nil && len(data) > 0 {
			files = append(files, filePart{field: "input_images", filename: filename, data: data})
			urlFilesAppended++
		}
	}
	if arr, ok := fields["input_image"]; ok {
		for _, u := range arr { appendURL(u) }
	}
	if arr, ok := fields["input_images"]; ok {
		for _, u := range arr { appendURL(u) }
	}

	// 决策 SKU（严格）：若秒数不合法将返回空 sku
	// Debug: 记录从 URL 追加的文件数量
	logger.SysDebug(fmt.Sprintf("sutui.rewrite debug -> url_files_appended=%d", urlFilesAppended))

	sku, slotSec := pickSutuiSKU(originalModel, size, seconds)
	if strings.TrimSpace(sku) == "" || slotSec == 0 {
		return nil, "", fmt.Errorf("invalid seconds for sutui: allowed 10/15 for sora-2, 15/25 for sora-2-pro")
	}

	// 重建 multipart
	var buf bytes.Buffer
	builder := p.Requester.CreateFormBuilder(&buf)

	// 回写文本字段：覆盖 model/seconds，其它保持
	wroteModel := false
	wroteSeconds := false
	for k, vals := range fields {
		switch strings.ToLower(k) {
		case "model":
			if err := builder.WriteField(k, sku); err != nil {
				return nil, "", err
			}
			wroteModel = true
		case "seconds":
			if slotSec > 0 {
				if err := builder.WriteField(k, strconv.Itoa(slotSec)); err != nil {
					return nil, "", err
				}
				wroteSeconds = true
			} else {
				// 使用严格时长
				if err := builder.WriteField(k, strconv.Itoa(slotSec)); err != nil {
					return nil, "", err
				}
				wroteSeconds = true
			}
		default:
			for _, v := range vals {
				if err := builder.WriteField(k, v); err != nil {
					return nil, "", err
				}
			}
		}
	}
	if !wroteModel {
		if err := builder.WriteField("model", sku); err != nil {
			return nil, "", err
		}
	}
	if !wroteSeconds && slotSec > 0 {
		if err := builder.WriteField("seconds", strconv.Itoa(slotSec)); err != nil {
			return nil, "", err
		}
	}

	// 文件字段
	for _, fp := range files {
		if err := builder.CreateFormFileReader(fp.field, bytes.NewReader(fp.data), fp.filename); err != nil {
			return nil, "", err
		}
	}

	if err := builder.Close(); err != nil {
		return nil, "", err
	}
	return &buf, builder.FormDataContentType(), nil
}

// validateMultipartSecondsEzlinkAI 仅校验 multipart 中 seconds 是否存在且为 4/8/12，不做改写
func (p *OpenAIProvider) validateMultipartSecondsEzlinkAI(raw []byte, contentType string) error {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil || !strings.HasPrefix(strings.ToLower(mediaType), "multipart/") {
		return fmt.Errorf("invalid multipart content-type")
	}
	boundary := params["boundary"]
	if strings.TrimSpace(boundary) == "" {
		return fmt.Errorf("missing boundary")
	}
	mr := multipart.NewReader(bytes.NewReader(raw), boundary)
	found := false
	for {
		part, e := mr.NextPart()
		if e == io.EOF {
			break
		}
		if e != nil {
			return e
		}
		name := part.FormName()
		if name == "" {
			_ = part.Close()
			continue
		}
		if part.FileName() != "" {
			_ = part.Close()
			continue
		}
		b, e2 := io.ReadAll(part)
		_ = part.Close()
		if e2 != nil {
			return e2
		}
		if strings.EqualFold(name, "seconds") {
			found = true
			v, err := strconv.Atoi(strings.TrimSpace(string(b)))
			if err != nil {
				return fmt.Errorf("seconds must be numeric and one of 4, 8, 12 for ezlinkai")
			}
			switch v {
			case 4, 8, 12:
			default:
				return fmt.Errorf("seconds must be one of 4, 8, 12 for ezlinkai")
			}
		}
	}
	if !found {
		return fmt.Errorf("missing seconds; ezlinkai requires seconds to be one of 4, 8, 12")
	}
	return nil
}

// pickSutuiSKU 根据官方模型 + 分辨率 + 秒数（宽松策略）映射到 sutui 的 sku 与秒数档位
// 规则：
// - 方向：由 size 推断，缺省 portrait
// - sora-2（SD）：<=10s -> 基础；>10s -> 15s
// - sora-2-pro（HD）：<=15s -> 15s；>15s -> 25s
func pickSutuiSKU(originalModel string, size string, seconds int) (sku string, slotSec int) {
	// 方向
	orientation := "portrait"
	s := strings.ReplaceAll(strings.ToLower(strings.TrimSpace(size)), " ", "")
	if strings.Contains(s, "x") {
		parts := strings.Split(s, "x")
		if len(parts) == 2 {
			w, _ := strconv.Atoi(parts[0])
			h, _ := strconv.Atoi(parts[1])
			if w >= h {
				orientation = "landscape"
			}
		}
	}
	// 严格校验允许的秒数：
	// - sora-2: 10 或 15
	// - sora-2-pro: 15 或 25
	if strings.EqualFold(strings.TrimSpace(originalModel), "sora-2-pro") {
		if !(seconds == 15 || seconds == 25) {
			return "", 0
		}
	} else {
		if !(seconds == 10 || seconds == 15) {
			return "", 0
		}
	}
	isPro := strings.EqualFold(strings.TrimSpace(originalModel), "sora-2-pro")

	base := "sora_video2"
	// 清晰度/时长档位
	if isPro {
		// HD 档
		if orientation == "portrait" {
			if seconds == 15 {
				return base + "-portrait-hd-15s", 15
			}
			return base + "-portrait-hd-25s", 25
		} else {
			if seconds == 15 {
				return base + "-landscape-hd-15s", 15
			}
			return base + "-landscape-hd-25s", 25
		}
	}
	// SD 档（sora-2）
	if orientation == "portrait" {
		if seconds == 10 {
			return base + "-portrait", 10
		}
		return base + "-portrait-15s", 15
	}
	if seconds == 10 {
		return base + "-landscape", 10
	}
	return base + "-landscape-15s", 15
}

// fetchImageForSutui 下载 URL 图片到内存并返回数据与建议文件名（带扩展名）。仅支持 http/https。
func fetchImageForSutui(u string, maxSize int) ([]byte, string, error) {
    if !(strings.HasPrefix(strings.ToLower(u), "http://") || strings.HasPrefix(strings.ToLower(u), "https://")) {
        return nil, "", fmt.Errorf("unsupported image url")
    }
    client := requester.HTTPClient
    if client == nil {
        client = http.DefaultClient
    }
    req, err := http.NewRequest(http.MethodGet, u, nil)
    if err != nil {
        return nil, "", err
    }
    // 设置较小的超时保护（如全局未设置）
    // 注意：requester.HTTPClient 可能已配置全局超时/代理
    resp, err := client.Do(req)
    if err != nil {
        return nil, "", err
    }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return nil, "", fmt.Errorf("bad status: %d", resp.StatusCode)
    }
    // 限流读取
    lr := io.LimitReader(resp.Body, int64(maxSize)+1)
    data, err := io.ReadAll(lr)
    if err != nil {
        return nil, "", err
    }
    if len(data) == 0 || len(data) > maxSize {
        return nil, "", fmt.Errorf("image too large or empty")
    }
    // 推断扩展名
    ext := ".jpg"
    ct := strings.ToLower(strings.TrimSpace(resp.Header.Get("Content-Type")))
    switch {
    case strings.Contains(ct, "image/png"):
        ext = ".png"
    case strings.Contains(ct, "image/webp"):
        ext = ".webp"
    case strings.Contains(ct, "image/jpeg"), strings.Contains(ct, "image/jpg"):
        ext = ".jpg"
    default:
        // 尝试从 URL 路径推断
        if u2, err := url.Parse(u); err == nil {
            if e := strings.ToLower(filepath.Ext(u2.Path)); e == ".png" || e == ".webp" || e == ".jpg" || e == ".jpeg" {
                ext = e
            }
        }
    }
    // 生成文件名
    filename := utils.GetUUID() + ext
    return data, filename, nil
}

func (p *OpenAIProvider) retrieveSutuiVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	fullURL := p.buildOfficialVideoURL(videoID, "")
	headers := p.GetRequestHeaders()

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}

	respObj := &sutuiRetrieveResponse{}
	_, errWith := p.Requester.SendRequest(req, respObj, false)
	if errWith != nil {
		return nil, errWith
	}

	job := &types.VideoJob{
		ID:          respObj.ID,
		Object:      respObj.Object,
		CreatedAt:   respObj.CreatedAt,
		CompletedAt: respObj.CompletedAt,
		Status:      respObj.Status,
		Model:       respObj.Model,
		Progress:    anyToFloat(respObj.Progress),
		Seconds:     respObj.Seconds,
		Size:        respObj.Size,
		Quality:     "standard",
	}
	if job.Object == "" {
		job.Object = "video"
	}
	// 覆盖非官方分辨率：若上游返回 720x720，按模型后缀推断方向并赋值官方分辨率
	if strings.TrimSpace(job.Size) == "720x720" {
		m := strings.ToLower(strings.TrimSpace(respObj.Model))
		if strings.Contains(m, "-landscape") {
			job.Size = "1280x720"
		} else if strings.Contains(m, "-portrait") {
			job.Size = "720x1280"
		}
	}

	if job.Result == nil && strings.TrimSpace(respObj.VideoURL) != "" {
		job.Result = &types.VideoJobResult{VideoURL: respObj.VideoURL}
	}
	return job, nil
}

func (p *OpenAIProvider) downloadSutuiVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	if strings.TrimSpace(variant) != "" && strings.ToLower(variant) != "video" {
		return nil, common.StringErrorWrapperLocal("variant not supported for Sutui channel", "unsupported_variant", http.StatusNotImplemented)
	}

	job, errWith := p.retrieveSutuiVideo(videoID)
	if errWith != nil {
		return nil, errWith
	}
	if job == nil || job.Result == nil || strings.TrimSpace(job.Result.VideoURL) == "" {
		return nil, common.StringErrorWrapperLocal("video url not ready", "video_not_ready", http.StatusBadRequest)
	}

	// 直接下载外链
	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	req, err := p.Requester.NewRequest(http.MethodGet, job.Result.VideoURL, p.Requester.WithHeader(map[string]string{}))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	resp, errWithCode := p.Requester.SendRequest(req, nil, true)
	if errWithCode != nil {
		return nil, errWithCode
	}
	return resp, nil
}

func (p *OpenAIProvider) remixApimartVideo(videoID string, prompt string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	modelName := strings.TrimSpace(p.GetOriginalModel())
	if modelName == "" {
		modelName = "sora-2"
	}
	body := map[string]any{
		"prompt": prompt,
		"model":  modelName,
	}
	headers := p.GetRequestHeaders()
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"
	fullURL := fmt.Sprintf("%s/v1/videos/%s/remix", strings.TrimSuffix(p.GetBaseURL(), "/"), url.PathEscape(videoID))
	req, err := p.Requester.NewRequest(http.MethodPost, fullURL, p.Requester.WithBody(body), p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	resp := &apimartCreateResponse{}
	if _, errWith := p.Requester.SendRequest(req, resp, false); errWith != nil {
		return nil, errWith
	}
	if resp.Code != http.StatusOK {
		return nil, apimartErrorToOpenAI(resp.Error, resp.Code)
	}
	if len(resp.Data) == 0 || strings.TrimSpace(resp.Data[0].TaskID) == "" {
		return nil, common.StringErrorWrapperLocal("invalid upstream response", "bad_upstream", http.StatusBadGateway)
	}
	job := &types.VideoJob{
		ID:                 resp.Data[0].TaskID,
		Object:             "video",
		Model:              modelName,
		Status:             mapApimartStatus(resp.Data[0].Status),
		Progress:           0,
		Quality:            "standard",
		CreatedAt:          time.Now().Unix(),
		Prompt:             prompt,
		RemixedFromVideoID: videoID,
	}
	if job.Status == "" {
		job.Status = "queued"
	}
	return job, nil
}

func (p *OpenAIProvider) retrieveApimartVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	data, errWith := p.getApimartTask(videoID)
	if errWith != nil {
		return nil, errWith
	}
	job := p.apimartTaskDataToVideoJob(videoID, data)
	return job, nil
}

func (p *OpenAIProvider) downloadApimartVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	data, errWith := p.getApimartTask(videoID)
	if errWith != nil {
		return nil, errWith
	}
	if data.Result == nil || len(data.Result.Videos) == 0 {
		return nil, common.StringErrorWrapperLocal("video not ready", "video_not_ready", http.StatusAccepted)
	}
	video := selectApimartVideo(data.Result.Videos, variant)
	urlStr := ""
	if video != nil {
		urlStr = video.URL.String()
	}
	if video == nil || strings.TrimSpace(urlStr) == "" {
		return nil, common.StringErrorWrapperLocal("variant not available", "unsupported_variant", http.StatusNotImplemented)
	}

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	req, err := p.Requester.NewRequest(http.MethodGet, urlStr, p.Requester.WithHeader(map[string]string{}))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	resp, errWith := p.Requester.SendRequest(req, nil, true)
	if errWith != nil {
		return nil, errWith
	}
	return resp, nil
}

func (p *OpenAIProvider) getApimartTask(videoID string) (*apimartTaskData, *types.OpenAIErrorWithStatusCode) {
	fullURL := strings.TrimSuffix(p.GetBaseURL(), "/") + "/v1/tasks/" + url.PathEscape(videoID)
	headers := p.GetRequestHeaders()
	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(headers))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}

	originalOpenAIFlag := p.Requester.IsOpenAI
	p.Requester.IsOpenAI = false
	defer func() { p.Requester.IsOpenAI = originalOpenAIFlag }()

	resp := &apimartTaskResponse{}
	if _, errWith := p.Requester.SendRequest(req, resp, false); errWith != nil {
		return nil, errWith
	}
	if resp.Code != http.StatusOK {
		return nil, apimartErrorToOpenAI(resp.Error, resp.Code)
	}
	if resp.Data == nil {
		return nil, common.StringErrorWrapperLocal("missing task data", "bad_upstream", http.StatusBadGateway)
	}
	return resp.Data, nil
}

func (p *OpenAIProvider) apimartTaskDataToVideoJob(fallbackID string, data *apimartTaskData) *types.VideoJob {
	job := &types.VideoJob{
		ID:          data.ID,
		Object:      "video",
		Status:      mapApimartStatus(data.Status),
		Progress:    anyToFloat(data.Progress),
		CreatedAt:   data.Created,
		CompletedAt: data.Completed,
		Model:       data.Model,
		Quality:     "standard",
	}
	if job.ID == "" {
		job.ID = fallbackID
	}
	if job.Model == "" {
		job.Model = p.GetOriginalModel()
	}
	if job.Status == "" {
		job.Status = "in_progress"
	}
	if data.TaskInfo != nil {
		job.Metadata = data.TaskInfo
	}
	if data.Error != nil {
		job.Error = &types.VideoJobError{
			Code:    data.Error.Code,
			Message: data.Error.Message,
		}
	}
	if data.Result != nil && len(data.Result.Videos) > 0 {
		video := data.Result.Videos[0]
		videoURL := video.URL.String()
		job.Result = &types.VideoJobResult{
			VideoURL:      videoURL,
			ThumbnailURL:  video.Cover,
			Variant:       video.Variant,
			DownloadURL:   videoURL,
			AdditionalRaw: data.Result,
		}
		if video.ExpiresAt > 0 {
			job.ExpiresAt = video.ExpiresAt
		}
		if video.Seconds > 0 && job.Seconds == 0 {
			job.Seconds = video.Seconds
		}
		if video.Resolution != "" && job.Size == "" {
			job.Size = video.Resolution
		}
	}
	return job
}

func selectApimartVideo(videos []apimartTaskVideo, variant string) *apimartTaskVideo {
	if len(videos) == 0 {
		return nil
	}
	cleanVariant := strings.ToLower(strings.TrimSpace(variant))
	if cleanVariant == "" || cleanVariant == "video" || cleanVariant == "mp4" {
		return &videos[0]
	}
	for i := range videos {
		if strings.ToLower(strings.TrimSpace(videos[i].Variant)) == cleanVariant ||
			strings.ToLower(strings.TrimSpace(videos[i].Quality)) == cleanVariant {
			return &videos[i]
		}
	}
	return nil
}

func buildApimartAspectRatio(size string) string {
	info := normalizeSoraSize(size)
	if info.Orientation == "portrait" {
		return "9:16"
	}
	if info.Orientation == "landscape" {
		return "16:9"
	}
	return ""
}

// normalizeApimartSeconds 将对外暴露的 OpenAI seconds 归一化到 apimart 支持的档位
// 规则：
// - sora-2：<=10 => 10；>10 => 15；缺省或非法 -> 10
// - sora-2-pro：<=15 => 15；>15 => 25；缺省或非法 -> 15
// - 其他模型名：按 sora-2 处理（更保守）
func normalizeApimartSeconds(model string, seconds int) int {
    m := strings.ToLower(strings.TrimSpace(model))
    if m == "sora-2-pro" {
        if seconds <= 0 { return 15 }
        if seconds <= 15 { return 15 }
        return 25
    }
    // 默认 sora-2
    if seconds <= 0 { return 10 }
    if seconds <= 10 { return 10 }
    return 15
}

func mapApimartStatus(status string) string {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "submitted", "pending":
		return "queued"
	case "processing":
		return "in_progress"
	case "completed", "success", "succeeded":
		return "completed"
	case "failed", "error":
		return "failed"
	case "cancelled", "canceled":
		return "failed"
	default:
		return status
	}
}

func apimartErrorToOpenAI(err *apimartError, code int) *types.OpenAIErrorWithStatusCode {
	if err == nil {
		err = &apimartError{
			Code:    code,
			Message: "upstream error",
			Type:    "upstream_error",
		}
	}
	statusCode := code
	if statusCode < 400 || statusCode > 599 {
		statusCode = http.StatusBadGateway
	}
	return &types.OpenAIErrorWithStatusCode{
		OpenAIError: types.OpenAIError{
			Message: err.Message,
			Code:    fmt.Sprintf("%d", err.Code),
			Type:    err.Type,
		},
		StatusCode: statusCode,
		LocalError: true,
	}
}

func normalizeSoraSize(size string) soraSizeInfo {
	value := strings.ToLower(strings.TrimSpace(size))
	value = strings.ReplaceAll(value, " ", "")

	if value == "" {
		return soraSizeInfo{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	}

	if strings.Contains(value, "x") {
		parts := strings.Split(value, "x")
		if len(parts) == 2 {
			w, _ := strconv.Atoi(parts[0])
			h, _ := strconv.Atoi(parts[1])
			orientation := "landscape"
			if h > w {
				orientation = "portrait"
			}
			resolution := fmt.Sprintf("%dx%d", w, h)
			return soraSizeInfo{
				Resolution:  resolution,
				Orientation: orientation,
				SizeLabel:   mapResolutionToSizeLabel(resolution),
			}
		}
	}

	switch value {
	case "landscape":
		return soraSizeInfo{
			Resolution:  "1280x720",
			Orientation: "landscape",
			SizeLabel:   "small",
		}
	case "portrait":
		return soraSizeInfo{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	default:
		return soraSizeInfo{
			Resolution:  "720x1280",
			Orientation: "portrait",
			SizeLabel:   "small",
		}
	}
}

func mapResolutionToSizeLabel(resolution string) string {
	switch resolution {
	case "1280x720", "720x1280":
		return "small"
	case "1792x1024", "1024x1792":
		return "large"
	default:
		return "small"
	}
}

func formatResolutionByOrientation(size string, orientation string) string {
	value := strings.ToLower(strings.TrimSpace(size))
	switch value {
	case "small":
		if orientation == "portrait" {
			return "720x1280"
		}
		return "1280x720"
	case "medium":
		if orientation == "portrait" {
			return "1024x1536"
		}
		return "1536x1024"
	case "large":
		if orientation == "portrait" {
			return "1024x1792"
		}
		return "1792x1024"
	default:
		return value
	}
}

func mapMountSeaStatus(status string) (string, float64) {
	switch strings.ToLower(status) {
	case "ready", "completed", "success":
		return "completed", 100
	case "failed", "error":
		return "failed", 0
	case "processing", "in_progress":
		return "in_progress", 66
	case "queue", "queued":
		return "queued", 0
	default:
		return "in_progress", 0
	}
}

func parseMountSeaTime(value string) (int64, error) {
	if value == "" {
		return 0, fmt.Errorf("empty time")
	}

	if ts, err := strconv.ParseInt(value, 10, 64); err == nil {
		// assume seconds
		if ts > 1e12 {
			return ts / 1000, nil
		}
		return ts, nil
	}

	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02T15:04:05.000Z07:00",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t.Unix(), nil
		}
	}

	return 0, fmt.Errorf("unsupported time format: %s", value)
}
