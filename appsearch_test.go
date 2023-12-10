package appsearch

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockHTTPClient struct{}

// Call simulates making an HTTP request and sends a mock response to the channel
func (m MockHTTPClient) Call(method, url, apiKey string, body io.Reader, ch chan<- *http.Response) {
	// Simulate response based on your test requirements
	// For example, you can create a mock response like this:
	mockResponse := &http.Response{
		Status:     "OK",
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"message": "Schema created successfully"}`)),
	}
	ch <- mockResponse
}
func TestCreateEngine(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Verify the request URL
		expectedURL := "/engines"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.Path)
		}

		// Verify the request body
		expectedBody := `{"name":"TestEngine"}`
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		// Return a sample response
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		w.WriteHeader(response.StatusCode)
	}))

	// Close the server when done
	defer server.Close()

	// Create an instance of AppSearch
	appSearch := AppSearch{
		Url:    server.URL,
		ApiKey: "test-api-key",
	}

	// Call the CreateEngine function
	response := appSearch.CreateEngine("TestEngine")

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, response.StatusCode)
	}
}
func TestIndexDocument(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Verify the request URL
		expectedURL := "/engines/TestEngine/documents"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.Path)
		}

		// Verify the request body
		expectedBody := "test body"
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		// Return a sample response
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		w.WriteHeader(response.StatusCode)
	}))

	// Close the server when done
	defer server.Close()

	// Create an instance of AppSearch
	appSearch := AppSearch{
		Url:        server.URL,
		ApiKey:     "test-api-key",
		EngineName: "TestEngine",
	}

	// Create a test body
	testBody := bytes.NewBufferString("test body")

	// Call the IndexDocument function
	response := appSearch.IndexDocument(testBody)

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, response.StatusCode)
	}
}
func TestListDocument(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Verify the request URL
		expectedURL := "/engines/TestEngine/documents/list"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.Path)
		}

		// Verify the request headers
		expectedAPIKey := "test-api-key"
		apiKey := r.Header.Get("Authorization")
		if apiKey != expectedAPIKey {
			t.Errorf("Expected API key %s, got %s", expectedAPIKey, apiKey)
		}

		// Return a sample response
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		w.WriteHeader(response.StatusCode)
	}))

	// Close the server when done
	defer server.Close()

	// Create an instance of AppSearch
	appSearch := AppSearch{
		Url:        server.URL,
		ApiKey:     "test-api-key",
		EngineName: "TestEngine",
	}

	// Call the ListDocument function
	response := appSearch.ListDocument(nil)

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, response.StatusCode)
	}
}
func TestSearch(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Verify the request URL
		expectedURL := "/engines/TestEngine/search"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.Path)
		}

		// Verify the request headers
		expectedAPIKey := "Bearer test-api-key"
		apiKey := r.Header.Get("Authorization")
		if apiKey != expectedAPIKey {
			t.Errorf("Expected API key %s, got %s", expectedAPIKey, apiKey)
		}

		// Verify the request body
		expectedBody := "test body"
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		// Return a sample response
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		w.WriteHeader(response.StatusCode)
	}))

	// Close the server when done
	defer server.Close()

	// Create an instance of AppSearch
	appSearch := AppSearch{
		Url:        server.URL,
		ApiKey:     "test-api-key",
		EngineName: "TestEngine",
	}

	// Create a test body
	testBody := bytes.NewBufferString("test body")

	// Call the Search function
	response := appSearch.Search(testBody)

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, response.StatusCode)
	}
}
func TestFindIds(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Verify the request URL
		expectedURL := "/engines/TestEngine/documents"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.Path)
		}

		// Verify the request headers
		expectedAPIKey := "Bearer test-api-key"
		apiKey := r.Header.Get("Authorization")
		if apiKey != expectedAPIKey {
			t.Errorf("Expected API key %s, got %s", expectedAPIKey, apiKey)
		}

		// Verify the request body
		expectedBody := `{"ids": [123]}`
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		// Return a sample response
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		w.WriteHeader(response.StatusCode)
	}))

	// Close the server when done
	defer server.Close()

	// Create an instance of AppSearch
	appSearch := AppSearch{
		Url:        server.URL,
		ApiKey:     "test-api-key",
		EngineName: "TestEngine",
	}

	// Call the FindIds function
	response := appSearch.FindIds("123")

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, response.StatusCode)
	}
}
func TestSuggestions(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Verify the request URL
		expectedURL := "/engines/TestEngine/query_suggestion"
		if r.URL.Path != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.Path)
		}

		// Verify the request headers
		expectedAPIKey := "Bearer test-api-key"
		apiKey := r.Header.Get("Authorization")
		if apiKey != expectedAPIKey {
			t.Errorf("Expected API key %s, got %s", expectedAPIKey, apiKey)
		}

		// Verify the request body
		expectedBody := `{"query": "test query"}`
		body, _ := io.ReadAll(r.Body)
		if string(body) != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, string(body))
		}

		// Return a sample response
		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		}
		w.WriteHeader(response.StatusCode)
	}))

	// Close the server when done
	defer server.Close()

	// Create an instance of AppSearch
	appSearch := AppSearch{
		Url:        server.URL,
		ApiKey:     "test-api-key",
		EngineName: "TestEngine",
	}

	// Call the Suggestions function
	response := appSearch.Suggestions("test query")

	// Verify the response status code
	expectedStatusCode := http.StatusOK
	if response.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, response.StatusCode)
	}
}
