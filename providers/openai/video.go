package openai

import (
	"fmt"
	"net/http"
	"net/url"
	"one-api/common"
	"one-api/common/config"
	"one-api/types"
	"strconv"
	"strings"
	"time"
)

const (
	soraVendorOfficial = "official"
	soraVendorMountSea = "mountsea"
)

type openAIVideoResponse struct {
	types.VideoJob
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

func (p *OpenAIProvider) CreateVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	// 统一遵循 OpenAI 标准视频路由：/v1/videos
	return p.createOfficialVideo(request)
}

func (p *OpenAIProvider) RetrieveVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	// 统一遵循 OpenAI 标准视频路由：/v1/videos
	return p.retrieveOfficialVideo(videoID)
}

func (p *OpenAIProvider) DownloadVideo(videoID string, variant string) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	// 统一遵循 OpenAI 标准视频路由：/v1/videos
	return p.downloadOfficialVideo(videoID, variant)
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

	return soraVendorOfficial
}

func (p *OpenAIProvider) createOfficialVideo(request *types.VideoCreateRequest) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	reqCopy := *request
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

	return &job, nil
}

func (p *OpenAIProvider) retrieveOfficialVideo(videoID string) (*types.VideoJob, *types.OpenAIErrorWithStatusCode) {
	fullURL := p.buildOfficialVideoURL(videoID, "")
	req, err := p.Requester.NewRequest(http.MethodGet, fullURL, p.Requester.WithHeader(p.GetRequestHeaders()))
	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}
	defer req.Body.Close()

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
		duration = 8
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

func normalizeSoraSize(size string) soraSizeInfo {
	value := strings.ToLower(strings.TrimSpace(size))
	value = strings.ReplaceAll(value, " ", "")

	if value == "" {
		return soraSizeInfo{
			Resolution:  "1280x720",
			Orientation: "landscape",
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
			Resolution:  "1280x720",
			Orientation: "landscape",
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
