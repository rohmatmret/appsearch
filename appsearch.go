package appsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	ApiKey     string
	Url        string
	EngineName string
)

type AppSearch struct {
	ApiKey     string
	Url        string
	EngineName string
}

type Client struct {
	ApiKey     string
	Url        string
	EngineName string
}

func NewAppSearch(ApiKey, Url, EngineName string) *AppSearch {
	return &AppSearch{
		ApiKey:     ApiKey,
		Url:        Url,
		EngineName: EngineName,
	}
}

func Connect() *AppSearch {
	return &AppSearch{
		ApiKey:     ApiKey,
		Url:        Url,
		EngineName: EngineName,
	}
}

func (c AppSearch) CreateEngine(EngineName string) *http.Response {
	values := map[string]string{"name": EngineName}
	body, err := json.Marshal(values)

	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	URL := fmt.Sprintf("%s/engines", c.Url)
	newHttp := NewHttpClient(http.Client{Timeout: time.Second * 5})
	ch := make(chan *http.Response)
	go newHttp.Call(http.MethodPost, URL, c.ApiKey, bytes.NewBuffer(body), ch)
	return <-ch
}

func (c AppSearch) ListEngine(page, limit string) *http.Response {
	URL := fmt.Sprintf("%s/engines", c.Url)
	payload := strings.NewReader(`{"page": ` + page + `, "limit": ` + limit + `}`)
	ch := make(chan *http.Response)
	go NewHttpClient(http.Client{Timeout: time.Second * 5}).Call(http.MethodGet, URL, c.ApiKey, payload, ch)
	return <-ch
}

func (c AppSearch) DeleteEngine(EngineName string) *http.Response {
	URL := fmt.Sprintf("%s/engines/%s", c.Url, EngineName)
	ch := make(chan *http.Response)
	go NewHttpClient(http.Client{Timeout: time.Second * 5}).Call(http.MethodDelete, URL, c.ApiKey, nil, ch)
	return <-ch
}

func (c AppSearch) IndexDocument(body io.Reader) *http.Response {
	URL := fmt.Sprintf("%s/engines/%s/documents", c.Url, c.EngineName)
	ch := make(chan *http.Response)
	go NewHttpClient(http.Client{Timeout: time.Second * 5}).Call(http.MethodPost, URL, c.ApiKey, body, ch)
	return <-ch
}

func (c AppSearch) ListDocument(page io.Reader) *http.Response {
	valid := c.Validate()
	if !valid {
		return &http.Response{
			Status:     "Invalid",
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Invalid data Configurations"}`)),
		}
	}

	newHttp := NewHttpClient(http.Client{Timeout: time.Second * 5})
	URL := fmt.Sprintf("%s/engines/%s/documents/list", c.Url, c.EngineName)
	ch := make(chan *http.Response)
	go newHttp.Call(http.MethodGet, URL, c.ApiKey, nil, ch)
	return <-ch
}

func (c AppSearch) FindIds(id string) *http.Response {
	valid := c.Validate()
	if !valid {
		return &http.Response{
			Status:     "Invalid",
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Invalid data Configurations"}`)),
		}
	}

	newHttp := NewHttpClient(http.Client{Timeout: time.Second * 5})
	URL := fmt.Sprintf("%s/engines/%s/documents/ids[%s]", c.Url, c.EngineName, id)
	ch := make(chan *http.Response)
	go newHttp.Call(http.MethodGet, URL, c.ApiKey, nil, ch)
	return <-ch
}

func (c AppSearch) Search(query io.Reader) *http.Response {
	valid := c.Validate()
	if !valid {
		return &http.Response{
			Status:     "Invalid",
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Invalid data Configurations"}`)),
		}
	}
	URL := fmt.Sprintf("%s/engines/%s/search", c.Url, c.EngineName)
	ch := make(chan *http.Response)
	go NewHttpClient(http.Client{Timeout: time.Second * 5}).Call(http.MethodPost, URL, c.ApiKey, query, ch)
	return <-ch
}

func (c AppSearch) Suggestions(query string) *http.Response {
	valid := c.Validate()
	if !valid {
		return &http.Response{
			Status:     "Invalid",
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Invalid data Configurations"}`)),
		}
	}
	payload := strings.NewReader(`{"query": "` + query + `"}`)
	URL := fmt.Sprintf("%s/engines/%s/query_suggestion", c.Url, c.EngineName)
	ch := make(chan *http.Response)
	go NewHttpClient(http.Client{Timeout: time.Second * 5}).Call(http.MethodPost, URL, c.ApiKey, payload, ch)
	return <-ch
}

func (c AppSearch) FilterDocument(filters io.Reader) *http.Response {
	URL := fmt.Sprintf("%s/engines/%s/documents/search", c.Url, c.EngineName)
	ch := make(chan *http.Response)
	go NewHttpClient(http.Client{Timeout: time.Second * 5}).Call(http.MethodPost, URL, c.ApiKey, filters, ch)
	return <-ch
}
