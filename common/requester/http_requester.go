package requester

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"one-api/common"
	"one-api/common/logger"
	"one-api/common/utils"
	"one-api/types"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpErrorHandler func(*http.Response) *types.OpenAIError

type HTTPRequester struct {
	// requestBuilder    utils.RequestBuilder
	CreateFormBuilder func(io.Writer) FormBuilder
	ErrorHandler      HttpErrorHandler
	proxyAddr         string
	Context           context.Context
	IsOpenAI          bool
}

// NewHTTPRequester 创建一个新的 HTTPRequester 实例。
// proxyAddr: 是代理服务器的地址。
// errorHandler: 是一个错误处理函数，它接收一个 *http.Response 参数并返回一个 *types.OpenAIErrorResponse。
// 如果 errorHandler 为 nil，那么会使用一个默认的错误处理函数。
func NewHTTPRequester(proxyAddr string, errorHandler HttpErrorHandler) *HTTPRequester {
	return &HTTPRequester{
		CreateFormBuilder: func(body io.Writer) FormBuilder {
			return NewFormBuilder(body)
		},
		ErrorHandler: errorHandler,
		proxyAddr:    proxyAddr,
		Context:      context.Background(),
		IsOpenAI:     true,
	}
}

type requestOptions struct {
	body   any
	header http.Header
}

type requestOption func(*requestOptions)

func (r *HTTPRequester) setProxy() context.Context {
	return utils.SetProxy(r.proxyAddr, r.Context)
}

// 创建请求
func (r *HTTPRequester) NewRequest(method, url string, setters ...requestOption) (*http.Request, error) {
    args := &requestOptions{
        body:   nil,
        header: make(http.Header),
    }
    for _, setter := range setters {
        setter(args)
    }
    req, err := utils.RequestBuilder(r.setProxy(), method, url, args.body, args.header)
    if err != nil {
        return nil, err
    }

    // Debug: 打印上游请求（方法、URL、Content-Type、部分请求体）
    // 仅在 debug 等级下有效，由上层 logger 控制输出
    ct := req.Header.Get("Content-Type")
    bodyPreview := ""
    switch b := args.body.(type) {
    case nil:
        // no body
    case io.Reader:
        // 避免读取消耗流，标记占位
        bodyPreview = "<stream>"
    case []byte:
        if strings.Contains(strings.ToLower(ct), "multipart/") {
            bodyPreview = fmt.Sprintf("<multipart len=%d>", len(b))
        } else {
            bodyPreview = truncateForLog(b, 4096)
        }
    default:
        if jb, e := json.Marshal(b); e == nil {
            bodyPreview = truncateForLog(jb, 4096)
        }
    }
    logger.SysDebug(fmt.Sprintf("up.req -> %s %s ct=%s len=%d body=%s", method, url, ct, req.ContentLength, bodyPreview))

    return req, nil
}

// 发送请求
func (r *HTTPRequester) SendRequest(req *http.Request, response any, outputResp bool) (*http.Response, *types.OpenAIErrorWithStatusCode) {
    resp, err := HTTPClient.Do(req)
    if err != nil {
        return nil, common.ErrorWrapper(err, "http_request_failed", http.StatusInternalServerError)
    }

	if !outputResp {
		defer resp.Body.Close()
	}

	// 处理响应
	if r.IsFailureStatusCode(resp) {
		return nil, HandleErrorResp(resp, r.ErrorHandler, r.IsOpenAI)
	}

    // 解析响应，并在 debug 下打印响应体（JSON 场景会截断）
    if response == nil {
        return resp, nil
    }

    if outputResp {
        var buf bytes.Buffer
        tee := io.TeeReader(resp.Body, &buf)
        err = DecodeResponse(tee, response)
        // 将响应体重新写入 resp.Body
        resp.Body = io.NopCloser(&buf)

        // Debug: 输出响应预览
        logger.SysDebug(fmt.Sprintf("up.resp <- status=%d ct=%s len≈%d body=%s", resp.StatusCode, resp.Header.Get("Content-Type"), buf.Len(), truncateForLog(buf.Bytes(), 4096)))
    } else {
        // 如果是 JSON，先读再解，便于打印
        ct := resp.Header.Get("Content-Type")
        if strings.Contains(strings.ToLower(ct), "application/json") {
            b, e := io.ReadAll(resp.Body)
            if e != nil {
                return nil, common.ErrorWrapper(e, "read_response_failed", http.StatusInternalServerError)
            }
            // 打印 debug
            logger.SysDebug(fmt.Sprintf("up.resp <- status=%d ct=%s len=%d body=%s", resp.StatusCode, ct, len(b), truncateForLog(b, 4096)))
            // decode
            if err = json.Unmarshal(b, response); err != nil {
                return nil, common.ErrorWrapper(err, "decode_response_failed", http.StatusInternalServerError)
            }
            // 重置 Body，供后续链路需要时读取
            resp.Body = io.NopCloser(bytes.NewReader(b))
        } else {
            // 非 JSON，直接 decode（无法打印体积大的二进制）
            err = json.NewDecoder(resp.Body).Decode(response)
        }
    }

    if err != nil {
        return nil, common.ErrorWrapper(err, "decode_response_failed", http.StatusInternalServerError)
    }

    return resp, nil
}

// 发送请求 RAW
func (r *HTTPRequester) SendRequestRaw(req *http.Request) (*http.Response, *types.OpenAIErrorWithStatusCode) {
	// 发送请求
	resp, err := HTTPClient.Do(req)
	if err != nil {
		return nil, common.ErrorWrapper(err, "http_request_failed", http.StatusInternalServerError)
	}

	// 处理响应
	if r.IsFailureStatusCode(resp) {
		return nil, HandleErrorResp(resp, r.ErrorHandler, r.IsOpenAI)
	}

	return resp, nil
}

// 获取流式响应
func RequestStream[T streamable](requester *HTTPRequester, resp *http.Response, handlerPrefix HandlerPrefix[T]) (*streamReader[T], *types.OpenAIErrorWithStatusCode) {
	// 如果返回的头是json格式 说明有错误
	// if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
	// 	return nil, HandleErrorResp(resp, requester.ErrorHandler, requester.IsOpenAI)
	// }

	stream := &streamReader[T]{
		reader:        bufio.NewReader(resp.Body),
		response:      resp,
		handlerPrefix: handlerPrefix,
		NoTrim:        false,

		DataChan: make(chan T),
		ErrChan:  make(chan error),
	}

	return stream, nil
}

func RequestNoTrimStream[T streamable](requester *HTTPRequester, resp *http.Response, handlerPrefix HandlerPrefix[T]) (*streamReader[T], *types.OpenAIErrorWithStatusCode) {
	stream, err := RequestStream(requester, resp, handlerPrefix)
	if err != nil {
		return nil, err
	}

	stream.NoTrim = true

	return stream, nil
}

// 设置请求体
func (r *HTTPRequester) WithBody(body any) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

// 设置请求头
func (r *HTTPRequester) WithHeader(header map[string]string) requestOption {
	return func(args *requestOptions) {
		for k, v := range header {
			args.header.Set(k, v)
		}
	}
}

// 设置Content-Type
func (r *HTTPRequester) WithContentType(contentType string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("Content-Type", contentType)
	}
}

// 判断是否为失败状态码
func (r *HTTPRequester) IsFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

// 处理错误响应
func HandleErrorResp(resp *http.Response, toOpenAIError HttpErrorHandler, isPrefix bool) *types.OpenAIErrorWithStatusCode {

	openAIErrorWithStatusCode := &types.OpenAIErrorWithStatusCode{
		StatusCode: resp.StatusCode,
		OpenAIError: types.OpenAIError{
			Message: "",
			Type:    "upstream_error",
			Code:    "bad_response_status_code",
			Param:   strconv.Itoa(resp.StatusCode),
		},
	}

	defer resp.Body.Close()

    if toOpenAIError != nil {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err == nil {
            resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
            errorResponse := toOpenAIError(resp)

			if errorResponse != nil && errorResponse.Message != "" {
				if strings.HasPrefix(errorResponse.Message, "当前分组") {
					openAIErrorWithStatusCode.StatusCode = http.StatusTooManyRequests
				}

				openAIErrorWithStatusCode.OpenAIError = *errorResponse
				if isPrefix {
					openAIErrorWithStatusCode.OpenAIError.Message = fmt.Sprintf("Provider API error: %s", openAIErrorWithStatusCode.OpenAIError.Message)
				}
			}

            // 如果 errorResponse 为 nil，并且响应体为JSON，则将响应体转换为字符串
            if errorResponse == nil && strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
                openAIErrorWithStatusCode.OpenAIError.Message = string(bodyBytes)
            }
            // Debug: 上游错误响应体
            logger.SysDebug(fmt.Sprintf("up.resp(error) <- status=%d ct=%s len=%d body=%s", resp.StatusCode, resp.Header.Get("Content-Type"), len(bodyBytes), truncateForLog(bodyBytes, 4096)))
        }
    }

	if openAIErrorWithStatusCode.OpenAIError.Message == "" {
		if isPrefix {
			openAIErrorWithStatusCode.OpenAIError.Message = fmt.Sprintf("Provider API error: bad response status code %d", resp.StatusCode)
		} else {
			openAIErrorWithStatusCode.OpenAIError.Message = fmt.Sprintf("bad response status code %d", resp.StatusCode)
		}
	}

	return openAIErrorWithStatusCode
}

// truncateForLog 将字节截断为安全的日志预览字符串（最多 n 字节），并去除换行以紧凑输出
func truncateForLog(b []byte, n int) string {
    if b == nil {
        return ""
    }
    if n <= 0 || len(b) <= n {
        return compact(string(b))
    }
    return compact(string(b[:n])) + "…(truncated)"
}

func compact(s string) string {
    // 简单压缩，避免多行撑爆日志
    s = strings.ReplaceAll(s, "\n", " ")
    s = strings.ReplaceAll(s, "\r", " ")
    return s
}

func SetEventStreamHeaders(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("X-Accel-Buffering", "no")
}

func GetJsonHeaders() map[string]string {
	return map[string]string{
		"Content-type": "application/json",
	}
}

type Stringer interface {
	GetString() *string
}

func DecodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return DecodeString(body, result)
	}

	if stringer, ok := v.(Stringer); ok {
		return DecodeString(body, stringer.GetString())
	}

	return json.NewDecoder(body).Decode(v)
}

func DecodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}
