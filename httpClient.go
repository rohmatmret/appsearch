package appsearch

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type HttpClient struct {
	Client http.Client
}

func NewHttpClient(client http.Client) *HttpClient {
	return &HttpClient{Client: client}
}

type HttpImplement interface {
	Call(method string, url, Apikey string, body io.Reader) *http.Response
}

func (h HttpClient) Call(method, url, Apikey string, body io.Reader, ch chan *http.Response) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", Apikey))
	return h.DoRequest(req, ch)
}

func (h HttpClient) DoRequest(req *http.Request, ch chan *http.Response) *http.Response {
	var mtx sync.Mutex
	mtx.Lock()
	resp, err := h.Client.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	mtx.Unlock()
	ch <- resp
	close(ch)
	return resp
}
