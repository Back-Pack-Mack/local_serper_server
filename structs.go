package main

import (
	"fmt"
	"net/http"
	"time"
)

type SearchResult struct {
	SearchParameters struct {
		Q      string `json:"q"`
		Type   string `json:"type,omitempty"`
		Engine string `json:"engine,omitempty"`
	} `json:"searchParameters"`
	AnswerBox       *AnswerBox      `json:"answerBox,omitempty"`
	KnowledgeGraph  *KnowledgeGraph `json:"knowledgeGraph,omitempty"`
	Organic         []OrganicResult `json:"organic,omitempty"`
	RelatedSearches []RelatedSearch `json:"relatedSearches,omitempty"`
	Images          []ImageResult   `json:"images,omitempty"`
	Videos          []VideoResult   `json:"videos,omitempty"`
	News            []NewsResult    `json:"news,omitempty"`
	Places          []PlaceResult   `json:"places,omitempty"`
	Error           string          `json:"error,omitempty"`
}

type AnswerBox struct {
	Title   string `json:"title"`
	Answer  string `json:"answer"`
	Snippet string `json:"snippet"`
	Link    string `json:"link"`
	Date    string `json:"date,omitempty"`
}

type KnowledgeGraph struct {
	Title             string              `json:"title"`
	Type              string              `json:"type"`
	Description       string              `json:"description"`
	DescriptionSource string              `json:"descriptionSource"`
	DescriptionLink   string              `json:"descriptionLink"`
	Attributes        []map[string]string `json:"attributes"`
	Image             string              `json:"image,omitempty"`
}

type OrganicResult struct {
	Title      string            `json:"title"`
	Link       string            `json:"link"`
	Snippet    string            `json:"snippet"`
	Position   int               `json:"position"`
	Date       string            `json:"date,omitempty"`
	Sitelinks  []SiteLink        `json:"sitelinks,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type SiteLink struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type RelatedSearch struct {
	Query string `json:"query"`
}

type ImageResult struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	ImageUrl string `json:"imageUrl"`
	Source   string `json:"source"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type VideoResult struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Source   string `json:"source"`
	Date     string `json:"date,omitempty"`
	Duration string `json:"duration,omitempty"`
}

type NewsResult struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Snippet  string `json:"snippet"`
	Date     string `json:"date"`
	Source   string `json:"source"`
	ImageUrl string `json:"imageUrl,omitempty"`
}

type PlaceResult struct {
	Title    string  `json:"title"`
	Address  string  `json:"address"`
	Rating   float64 `json:"rating"`
	Reviews  int     `json:"reviews"`
	Category string  `json:"category"`
	Phone    string  `json:"phone,omitempty"`
	Website  string  `json:"website,omitempty"`
}

type MockSerperClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	MockMode   bool
	Port       int
	Server     *http.Server
}

func NewMockSerperClient(apiKey string, port int) *MockSerperClient {
	if apiKey == "" {
		apiKey = "mock-api-key"
	}
	if port == 0 {
		port = 8080
	}

	return &MockSerperClient{
		APIKey:  apiKey,
		BaseURL: fmt.Sprintf("http://localhost:%d", port),
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		MockMode: true,
		Port:     port,
	}
}
