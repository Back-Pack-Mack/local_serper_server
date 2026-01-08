package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *MockSerperClient) Search(query string, options ...SearchOption) (*SearchResult, error) {
	if c.MockMode {
		return generateMockResults(query, "search"), nil
	}

	reqBody := map[string]interface{}{
		"q": query,
	}

	for _, opt := range options {
		opt(reqBody)
	}

	return c.makeRequest(reqBody)
}

func (c *MockSerperClient) SearchImages(query string, options ...SearchOption) (*SearchResult, error) {
	if c.MockMode {
		result := generateMockResults(query, "images")
		result.Organic = nil
		result.KnowledgeGraph = nil
		result.AnswerBox = nil
		return result, nil
	}

	reqBody := map[string]interface{}{
		"q":    query,
		"type": "images",
	}

	for _, opt := range options {
		opt(reqBody)
	}

	return c.makeRequest(reqBody)
}

func (c *MockSerperClient) SearchVideos(query string, options ...SearchOption) (*SearchResult, error) {
	if c.MockMode {
		result := generateMockResults(query, "videos")
		result.Organic = nil
		result.KnowledgeGraph = nil
		result.AnswerBox = nil
		return result, nil
	}

	reqBody := map[string]interface{}{
		"q":    query,
		"type": "videos",
	}

	for _, opt := range options {
		opt(reqBody)
	}

	return c.makeRequest(reqBody)
}

func (c *MockSerperClient) SearchNews(query string, options ...SearchOption) (*SearchResult, error) {
	if c.MockMode {
		result := generateMockResults(query, "news")
		result.Organic = nil
		result.KnowledgeGraph = nil
		result.AnswerBox = nil
		return result, nil
	}

	reqBody := map[string]interface{}{
		"q":    query,
		"type": "news",
	}

	for _, opt := range options {
		opt(reqBody)
	}

	return c.makeRequest(reqBody)
}

func (c *MockSerperClient) SearchPlaces(query string, options ...SearchOption) (*SearchResult, error) {
	if c.MockMode {
		result := generateMockResults(query, "places")
		result.Organic = nil
		result.KnowledgeGraph = nil
		result.AnswerBox = nil
		return result, nil
	}

	reqBody := map[string]interface{}{
		"q":    query,
		"type": "places",
	}

	for _, opt := range options {
		opt(reqBody)
	}

	return c.makeRequest(reqBody)
}

func (c *MockSerperClient) makeRequest(body map[string]interface{}) (*SearchResult, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://google.serper.dev/search", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("X-API-KEY", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result, nil
}

func (c *MockSerperClient) StartMockServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/search", mockHandler)
	mux.HandleFunc("/images", mockHandler)
	mux.HandleFunc("/videos", mockHandler)
	mux.HandleFunc("/news", mockHandler)
	mux.HandleFunc("/places", mockHandler)

	c.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Mock Serper API server started on http://localhost:%d\n", c.Port)
	return c.Server.ListenAndServe()
}

func (c *MockSerperClient) StopMockServer() error {
	if c.Server != nil {
		return c.Server.Close()
	}
	return nil
}
