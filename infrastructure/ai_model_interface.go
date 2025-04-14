package infrastructure

import "github.com/yesetoda/BlogMate/domain"

type Data struct {
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}
type DataXTitle struct {
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}
type DataXContent struct {
	Title string   `json:"title,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}
type DataXTags struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
type AIModel interface {
	RecommendBlogs(Data) ([]domain.BlogRecommendation, error)
	RecommendTitle(content string, tags []string) (string, error)
	RecommendContent(title string, tags []string) (string, error)
	RecommendTags(title string, content string) ([]string, error)
	Summarize(Data) (string, error)
	Validate(Data) error
	Chat(prompt string) (string, error)
}
