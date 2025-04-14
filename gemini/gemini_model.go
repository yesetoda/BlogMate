package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/yesetoda/BlogMate/domain"
	"github.com/yesetoda/BlogMate/infrastructure"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiModel struct {
	model   *genai.GenerativeModel
	prompts infrastructure.Prompts
}



func connectToGemini(apiKey, modelName string, ctx context.Context) (*genai.GenerativeModel, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return &genai.GenerativeModel{}, err
	}
	return client.GenerativeModel(modelName), nil
}

func NewGeminiModel(apiKey, modelName string, prompts infrastructure.Prompts) (*GeminiModel, error) {
	model, err := connectToGemini(apiKey, modelName, context.Background())
	if err != nil {
		return &GeminiModel{}, err
	}
	return &GeminiModel{model: model, prompts: prompts}, nil
}

func (g *GeminiModel) SendPrompt(prompt string) (string, error) {
	resp, err := g.model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("no response from the model")
	}
	candidate := resp.Candidates[0]
	responseText := fmt.Sprint(candidate.Content.Parts[0])
	return responseText, nil
}

func (g *GeminiModel) CheckPromptContent(content string) error {
	p := fmt.Sprintf(g.prompts.CheckPromptContent, content)
	resp, err := g.SendPrompt(p)
	if err != nil {
		return fmt.Errorf("getting response error: %v", err.Error())
	}
	if strings.ToLower(resp) == "no" {
		return fmt.Errorf("prompt not about blogging")
	}
	return nil
}

func (g *GeminiModel) Refine(content string) (string, error) {
	prompt := fmt.Sprintf(g.prompts.Refine, content)
	refinedContent, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return refinedContent, nil
}

func (g *GeminiModel) Validate(data infrastructure.Data) error {
	prompt := fmt.Sprintf(g.prompts.Validate, data.Content)
	validation, err := g.SendPrompt(prompt)
	if err != nil {
		return err
	}
	if validation != "yes" {
		return fmt.Errorf("%v", validation)
	}
	return nil
}

func (g *GeminiModel) RecommendBlogs(data infrastructure.Data) ([]domain.BlogRecommendation, error) {
	prompt := fmt.Sprintf(
		g.prompts.RecommendBlog,
		data.Title,
		data.Content,
		strings.Join(data.Tags, ", "),
	)

	resp, err := g.SendPrompt(prompt)
	if err != nil {
		return nil, err
	}

	// Clean response for JSON parsing
	resp = strings.TrimSpace(resp)
	resp = strings.Trim(resp, "```json")
	resp = strings.Trim(resp, "```")

	var blogs []domain.BlogRecommendation
	if err := json.Unmarshal([]byte(resp), &blogs); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Ensure exactly 5 recommendations
	if len(blogs) < 5 {
		return nil, fmt.Errorf("insufficient recommendations generated")
	}

	return blogs[:5], nil
}

func (g *GeminiModel) Summarize(data infrastructure.Data) (string, error) {
	blog := fmt.Sprintf("Title: %v, Content %v, Tags %v", data.Title, data.Content, data.Tags)
	prompt := fmt.Sprintf(g.prompts.Summarize, blog)
	summary, err := g.SendPrompt(prompt)
	return summary, err
}

func (g *GeminiModel) RecommendTitle(content string, tags []string) (string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendTitle, content, strings.Join(tags, ", "))
	if err := g.CheckPromptContent(prompt); err != nil {
		return "", err
	}
	resp, err := g.SendPrompt(prompt)
	titles := strings.Split(resp, "\n")
	if err != nil {
		return "", err
	}
	return titles[0], nil
}

func (g *GeminiModel) Chat(content string) (string, error) {
	if err := g.CheckPromptContent(content); err != nil {
		return "", err
	}
	resp, err := g.SendPrompt(content)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (g *GeminiModel) RecommendContent(title string, tags []string) (string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendContent, title, tags)
	if err := g.CheckPromptContent(prompt); err != nil {
		return "", err
	}
	recommendedContent, err := g.SendPrompt(prompt)
	if err != nil {
		return "", err
	}
	return recommendedContent, nil
}

func (g *GeminiModel) RecommendTags(title string, content string) ([]string, error) {
	prompt := fmt.Sprintf(g.prompts.RecommendContent, title, content)
	if err := g.CheckPromptContent(prompt); err != nil {
		return []string{}, err
	}
	t, err := g.SendPrompt(prompt)
	tags := strings.Split(t, ",")
	if err != nil {
		return []string{}, err
	}
	return tags, nil
}
