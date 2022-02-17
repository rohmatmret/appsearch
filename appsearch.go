package appsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func NewAppSearch(ApiKey, Url, EngineName string) *AppSearch {
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
		log.Fatal(err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	h := http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines", c.Url), bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}

func (c AppSearch) ListEngine(page io.Reader) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/engines", c.Url), page)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}

func (c AppSearch) DeleteEngine(EngineName string) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/engines/%s", c.Url, EngineName), nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}

func (c AppSearch) IndexDocument(body io.Reader) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/documents", c.Url, c.EngineName), body)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}

func (c AppSearch) ListDocument(page io.Reader) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/engines/%s/documents/list?page[size]=1&page[current]=1", c.Url, c.EngineName), page)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}

func (c AppSearch) FindIds(id string) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/engines/%s/documents/ids[%s]", c.Url, c.EngineName, id), nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()

	return resp
}

func (c AppSearch) Search(query string, page io.Reader) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/search", c.Url, c.EngineName), page)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	fmt.Println(req.URL)
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}

func (c AppSearch) Suggestions(query string) *http.Response {
	h := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/engines/%s/query_suggestion", c.Url, c.EngineName), nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	resp, err := h.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return &http.Response{
			Status:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer resp.Body.Close()
	return resp
}
