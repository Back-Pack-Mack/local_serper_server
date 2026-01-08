package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func generateMockResults(query, searchType string) *SearchResult {
	rand.Seed(time.Now().UnixNano())

	result := &SearchResult{
		SearchParameters: struct {
			Q      string `json:"q"`
			Type   string `json:"type,omitempty"`
			Engine string `json:"engine,omitempty"`
		}{
			Q:      query,
			Type:   searchType,
			Engine: "google",
		},
	}

	switch strings.ToLower(searchType) {
	case "search":
		generateSearchResults(result, query)
	case "images":
		generateImageResults(result, query)
	case "videos":
		generateVideoResults(result, query)
	case "news":
		generateNewsResults(result, query)
	case "places":
		generatePlaceResults(result, query)
	default:
		generateSearchResults(result, query)
	}

	return result
}

func generateSearchResults(result *SearchResult, query string) {
	if rand.Intn(100) > 30 {
		result.KnowledgeGraph = &KnowledgeGraph{
			Title:             fmt.Sprintf("关于 %s 的信息", query),
			Type:              "一般信息",
			Description:       fmt.Sprintf("%s 是一个广泛讨论的话题，涉及多个领域。", query),
			DescriptionSource: "维基百科",
			DescriptionLink:   "https://zh.wikipedia.org/wiki/" + url.QueryEscape(query),
			Attributes: []map[string]string{
				{"类别": "科技"},
				{"重要性": "高"},
				{"流行度": "上升中"},
			},
		}
	}

	result.Organic = []OrganicResult{
		{
			Title:    fmt.Sprintf("深入了解 %s 的完整指南", query),
			Link:     fmt.Sprintf("https://example.com/guide/%s", url.QueryEscape(query)),
			Snippet:  fmt.Sprintf("本文详细介绍了%s的概念、应用场景以及最佳实践。", query),
			Position: 1,
			Date:     time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
			Sitelinks: []SiteLink{
				{Title: fmt.Sprintf("%s 入门教程", query), Link: "https://example.com/tutorial"},
				{Title: fmt.Sprintf("%s 常见问题", query), Link: "https://example.com/faq"},
			},
		},
		{
			Title:    fmt.Sprintf("%s 的最新发展和趋势", query),
			Link:     "https://technews.com/latest/" + url.QueryEscape(query),
			Snippet:  "探索最新的技术进展和行业趋势，了解如何利用这些技术优化您的工作流程。",
			Position: 2,
			Date:     time.Now().AddDate(0, 0, -3).Format("2006-01-02"),
		},
		{
			Title:    "官方文档和技术规范",
			Link:     "https://docs.example.com/" + url.QueryEscape(query),
			Snippet:  "查看完整的API文档、使用示例和技术规范，帮助您更好地理解和应用相关技术。",
			Position: 3,
		},
		{
			Title:    "社区讨论和最佳实践",
			Link:     "https://community.example.com/tags/" + url.QueryEscape(query),
			Snippet:  "加入开发者社区，与其他专家交流经验，分享您在使用过程中遇到的问题和解决方案。",
			Position: 4,
		},
		{
			Title:    "视频教程和在线课程",
			Link:     "https://learning.example.com/courses/" + url.QueryEscape(query),
			Snippet:  "通过互动式学习路径，从基础到高级，系统性地掌握相关技能。",
			Position: 5,
		},
	}

	result.RelatedSearches = []RelatedSearch{
		{Query: query + " 教程"},
		{Query: query + " 入门指南"},
		{Query: query + " 最佳实践"},
		{Query: "如何学习 " + query},
		{Query: query + " vs 替代方案"},
	}

	if strings.Contains(query, "是什么") || strings.Contains(query, "什么是") {
		result.AnswerBox = &AnswerBox{
			Title:   fmt.Sprintf("%s 的定义", strings.TrimPrefix(query, "什么是")),
			Answer:  fmt.Sprintf("%s 是一种重要的技术概念，广泛应用于现代软件开发中。", strings.TrimPrefix(query, "什么是")),
			Snippet: "了解更多详细信息请参考官方文档和社区资源。",
			Link:    "https://zh.wikipedia.org/wiki/" + url.QueryEscape(query),
		}
	}
}

func generateImageResults(result *SearchResult, query string) {
	result.Images = []ImageResult{
		{
			Title:    fmt.Sprintf("%s 示意图", query),
			Link:     fmt.Sprintf("https://images.example.com/%s-1.jpg", url.QueryEscape(query)),
			ImageUrl: fmt.Sprintf("https://images.example.com/%s-1.jpg", url.QueryEscape(query)),
			Source:   "Example Images",
			Width:    800,
			Height:   600,
		},
		{
			Title:    fmt.Sprintf("%s 架构图", query),
			Link:     fmt.Sprintf("https://images.example.com/%s-2.jpg", url.QueryEscape(query)),
			ImageUrl: fmt.Sprintf("https://images.example.com/%s-2.jpg", url.QueryEscape(query)),
			Source:   "Tech Diagrams",
			Width:    1024,
			Height:   768,
		},
	}
}

func generateVideoResults(result *SearchResult, query string) {
	result.Videos = []VideoResult{
		{
			Title:    fmt.Sprintf("%s 入门教程", query),
			Link:     fmt.Sprintf("https://videos.example.com/%s-tutorial", url.QueryEscape(query)),
			Source:   "Tech Tutorials",
			Date:     "2024-01-15",
			Duration: "15:30",
		},
		{
			Title:    fmt.Sprintf("%s 高级技巧", query),
			Link:     fmt.Sprintf("https://videos.example.com/%s-advanced", url.QueryEscape(query)),
			Source:   "Dev Channel",
			Date:     "2024-01-10",
			Duration: "22:45",
		},
	}
}

func generateNewsResults(result *SearchResult, query string) {
	result.News = []NewsResult{
		{
			Title:   fmt.Sprintf("%s 最新发展动态", query),
			Link:    fmt.Sprintf("https://news.example.com/%s-latest", url.QueryEscape(query)),
			Snippet: "了解最新的技术进展和行业应用案例。",
			Date:    time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
			Source:  "Tech News Daily",
		},
		{
			Title:   fmt.Sprintf("专家解读 %s 的未来趋势", query),
			Link:    fmt.Sprintf("https://news.example.com/%s-trends", url.QueryEscape(query)),
			Snippet: "行业专家分享他们对未来发展的见解和预测。",
			Date:    time.Now().AddDate(0, 0, -2).Format("2006-01-02"),
			Source:  "Future Tech Review",
		},
	}
}

func generatePlaceResults(result *SearchResult, query string) {
	result.Places = []PlaceResult{
		{
			Title:    fmt.Sprintf("%s 技术公司", query),
			Address:  "上海市浦东新区张江高科技园区",
			Rating:   4.5,
			Reviews:  128,
			Category: "科技公司",
			Phone:    "021-12345678",
			Website:  "https://example.com",
		},
		{
			Title:    fmt.Sprintf("%s 研发中心", query),
			Address:  "北京市海淀区中关村",
			Rating:   4.7,
			Reviews:  256,
			Category: "研发机构",
		},
	}
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Q    string `json:"q"`
		Type string `json:"type,omitempty"`
		Num  int    `json:"num,omitempty"`
		Page int    `json:"page,omitempty"`
		Gl   string `json:"gl,omitempty"`
		Hl   string `json:"hl,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Q == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		req.Type = "search"
	}

	result := generateMockResults(req.Q, req.Type)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	json.NewEncoder(w).Encode(result)
}
