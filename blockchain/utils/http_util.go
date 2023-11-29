package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type HttpResponse struct {
	H http.Header
	B []byte
}

func (r *HttpResponse) Body() []byte {
	return r.B
}

func (r *HttpResponse) Header() http.Header {
	return r.H
}

func HttpPostOfJson(url string, obj interface{}, header map[string]string) (*HttpResponse, error) {
	j, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// 构造请求
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	// 添加header
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("panic: %v\n%s", err, string(debug.Stack()))
		}
	}()
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// 判断服务返回对象对不对 目前还没做正常的判断
	if response == nil {
		return nil, fmt.Errorf("response is nil")
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &HttpResponse{
		H: response.Header,
		B: body,
	}, nil
}

func HttpGetQuery(url string, params, header map[string]string) (*HttpResponse, error) {
	// 构建请求
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// 添加header
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}
	// 添加参数
	q := request.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		request.URL.RawQuery = q.Encode()
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("panic: %v\n%s", err, string(debug.Stack()))
		}
	}()
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// 判断服务返回对象对不对 目前还没做正常的判断
	if response == nil {
		return nil, fmt.Errorf("response is nil")
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &HttpResponse{
		H: response.Header,
		B: body,
	}, nil
}
