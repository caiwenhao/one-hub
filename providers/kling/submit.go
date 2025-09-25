package kling

import (
	"fmt"
	"net/http"

	"one-api/common"
	"one-api/types"
)

func (s *KlingProvider) Submit(class, action string, request *KlingTask) (data *KlingResponse[KlingTaskData], errWithCode *types.OpenAIErrorWithStatusCode) {
	submitUri := fmt.Sprintf(s.Generations, class, action)

	fullRequestURL := s.GetFullRequestURL(submitUri, "")
	headers := s.GetRequestHeaders()

	// 创建请求
	req, err := s.Requester.NewRequest(http.MethodPost, fullRequestURL, s.Requester.WithHeader(headers), s.Requester.WithBody(request))

	if err != nil {
		return nil, common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}

	data = &KlingResponse[KlingTaskData]{}
	_, errWithCode = s.Requester.SendRequest(req, data, false)

	return data, errWithCode
}

// CallCustomPath 调用自定义路径，便于多模态编辑等扩展接口复用
func (s *KlingProvider) CallCustomPath(method, path string, payload any, resp any) (errWithCode *types.OpenAIErrorWithStatusCode) {
	fullRequestURL := s.GetFullRequestURL(path, "")
	headers := s.GetRequestHeaders()

	var (
		req *http.Request
		err error
	)

	if payload != nil {
		req, err = s.Requester.NewRequest(method, fullRequestURL, s.Requester.WithHeader(headers), s.Requester.WithBody(payload))
	} else {
		req, err = s.Requester.NewRequest(method, fullRequestURL, s.Requester.WithHeader(headers))
	}
	if err != nil {
		return common.ErrorWrapper(err, "new_request_failed", http.StatusInternalServerError)
	}

	_, errWithCode = s.Requester.SendRequest(req, resp, false)
	return errWithCode
}
