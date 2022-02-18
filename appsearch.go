package appsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

	h := http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines", c.Url), bytes.NewBuffer(body))
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return resp
}

func (c AppSearch) ListEngine(page byte) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	body := bytes.NewBuffer([]byte(fmt.Sprintf("page[size]=1&page[current]=%d", page)))
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/engines", c.Url), body)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return resp
}

func (c AppSearch) DeleteEngine(EngineName string) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/engines/%s", c.Url, EngineName), nil)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return resp
}

func (c AppSearch) IndexDocument(body string) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/documents", c.Url, c.EngineName), nil)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return resp
}

type Page struct {
	Current int `json:"current"`
	Size    int `json:"size"`
}

type Result interface{}

func (c AppSearch) ListDocument(page io.Reader) *http.Response {
	valid := c.Validate()
	if !valid {
		return &http.Response{
			Status:     "Invalid",
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Invalid data Configurations"}`)),
		}
	}

	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/engines/%s/documents/list", c.Url, c.EngineName), page)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return resp
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

	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/engines/%s/documents/ids[%s]", c.Url, c.EngineName, id), nil)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return resp
}

func (c AppSearch) Search(query string, page io.Reader) *http.Response {
	valid := c.Validate()
	if !valid {
		return &http.Response{
			Status:     "Invalid",
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"error": "Invalid data Configurations"}`)),
		}
	}
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/search", c.Url, c.EngineName), page)
	if err != nil {
		fmt.Println("error new request", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return resp
}

func (c AppSearch) Suggestions(query string) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/query_suggestion", c.Url, c.EngineName), nil)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return resp
}

func (c AppSearch) FilterDocument(filters io.Reader) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/search", c.Url, c.EngineName), filters)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return resp
}
