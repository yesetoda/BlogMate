package controllers

import (
	"net/http"

	"github.com/yesetoda/BlogMate/infrastructure"

	"github.com/gin-gonic/gin"
)

// AIController handles AI-related operations
// @Description Controller for AI-powered recommendations and interactions
type AIController struct {
	model infrastructure.AIModel
}

// NewAIController creates a new AIController instance
func NewAIController(model infrastructure.AIModel) *AIController {
	return &AIController{model: model}
}

// ChatRequest represents the input for chat endpoint
type ChatRequest struct {
	Message string `json:"message,omitempty" example:"Hello, how are you?"`
}

// RecommendationResponse represents the response for recommendation endpoints
type RecommendationResponse struct {
	Titles   []string `json:"titles,omitempty"`
	Contents []string `json:"contents,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

// SummaryResponse represents the response for summarize endpoint
type SummaryResponse struct {
	Summary string `json:"summary" example:"This is a summary of the provided content."`
}

// ChatResponse represents the response for chat endpoint
type ChatResponse struct {
	Response string `json:"response" example:"I'm doing well, thank you for asking!"`
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error string `json:"error" example:"invalid data format"`
}

// Recommend godoc
// @Summary Get recommendations
// @Description Get AI-generated recommendations based on input data
// @Tags AI
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for recommendation"
// @Success 200 {object} interface{} "Successfully generated recommendations"
// @Failure 406 {object} ErrorResponse "Invalid input format"
// @Failure 500 {object} ErrorResponse "AI model error"
// @Router /ai/recommend [post]
func (c *AIController) RecommendBlog(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, ErrorResponse{Error: "invalid data format"})
		return
	}
	blogs, err := c.model.RecommendBlogs(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no response from the ai model"})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

// RecommendTitle godoc
// @Summary Get title recommendations
// @Description Get AI-generated title recommendations based on input data
// @Tags AI
// @Accept json
// @Produce json
// @Param input body infrastructure.DataXTitle true "Input data for recommendation"
// @Success 200 {object} RecommendationResponse "Successfully generated title recommendations"
// @Failure 406 {object} ErrorResponse "Invalid input format"
// @Failure 500 {object} ErrorResponse "AI model error"
// @Router /ai/recommend-title [post]
func (c *AIController) RecommendTitle(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, ErrorResponse{Error: "invalid data format"})
		return
	}
	titles, err := c.model.RecommendTitle(data.Content,data.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Titles": titles})
}

// RecommendContent godoc
// @Summary Get content recommendations
// @Description Get AI-generated content recommendations based on input data
// @Tags AI
// @Accept json
// @Produce json
// @Param input body infrastructure.DataXContent true "Input data for recommendation"
// @Success 200 {object} RecommendationResponse "Successfully generated content recommendations"
// @Failure 406 {object} ErrorResponse "Invalid input format"
// @Failure 500 {object} ErrorResponse "AI model error"
// @Router /ai/recommend-content [post]
func (c *AIController) RecommendContent(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, ErrorResponse{Error: "invalid data format"})
		return
	}
	contents, err := c.model.RecommendContent(data.Title,data.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"contents": contents})
}

// RecommendTags godoc
// @Summary Get tags recommendations
// @Description Get AI-generated tags recommendations based on input data
// @Tags AI
// @Accept json
// @Produce json
// @Param input body infrastructure.DataXTags true "Input data for recommendation"
// @Success 200 {object} RecommendationResponse "Successfully generated tags recommendations"
// @Failure 406 {object} ErrorResponse "Invalid input format"
// @Failure 500 {object} ErrorResponse "AI model error"
// @Router /ai/recommend-tags [post]
func (c *AIController) RecommendTags(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, ErrorResponse{Error: "invalid data format"})
		return
	}
	tags, err := c.model.RecommendTags(data.Title,data.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tags": tags})
}

// Summarize godoc
// @Summary Summarize content
// @Description Generate AI-powered summary of input content
// @Tags AI
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data to summarize"
// @Success 200 {object} SummaryResponse "Generated summary"
// @Failure 406 {object} ErrorResponse "Invalid input format"
// @Failure 500 {object} ErrorResponse "AI model error"
// @Router /ai/summarize [post]
func (c *AIController) Summarize(ctx *gin.Context) {
	data := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusNotAcceptable, ErrorResponse{Error: "invalid data format"})
		return
	}
	summary, err := c.model.Summarize(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "no response from the ai model"})
		return
	}
	ctx.JSON(http.StatusOK, SummaryResponse{Summary: summary})
}

// Chat godoc
// @Summary Chat with AI
// @Description Have a conversation with the AI model
// @Tags AI
// @Accept json
// @Produce json
// @Param input body ChatRequest true "Chat message"
// @Success 200 {object} ChatResponse "AI response"
// @Failure 406 {object} ErrorResponse "Invalid input format"
// @Failure 500 {object} ErrorResponse "AI model error"
// @Router /ai/chat [post]
func (c *AIController) Chat(ctx *gin.Context) {
	var request ChatRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusNotAcceptable, ErrorResponse{Error: "invalid data format"})
		return
	}
	response, err := c.model.Chat(request.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ChatResponse{Response: response})
}
