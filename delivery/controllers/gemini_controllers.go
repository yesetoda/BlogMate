package controllers

import (
	"net/http"

	"github.com/yesetoda/BlogMate/gemini"
	"github.com/yesetoda/BlogMate/infrastructure"

	"github.com/gin-gonic/gin"
)

type GeminiController struct {
	model *gemini.GeminiModel
}

// RecommendTitleController godoc
// @Summary Get title recommendations
// @Description Get AI-generated title recommendations based on content and tags
// @Tags AI Recommendations
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for title recommendation"
// @Success 200 {object} map[string]string "response: Recommended titles"
// @Failure 406 {object} map[string]string "error: Invalid input data"
// @Failure 500 {object} map[string]string "error: AI model error"
// @Router /recommend/title [post]
func (g *GeminiController) RecommendTitleController(ctx *gin.Context) {
	sampleBlog := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	resp, err := g.model.RecommendTitle(sampleBlog.Content, sampleBlog.Tags)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
}

// RecommendContent godoc
// @Summary Get content recommendations
// @Description Get AI-generated content recommendations based on title and tags
// @Tags AI Recommendations
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for content recommendation"
// @Success 200 {object} map[string]string "response: Recommended content"
// @Failure 406 {object} map[string]string "error: Invalid input data"
// @Failure 500 {object} map[string]string "error: AI model error"
// @Router /recommend/content [post]
func (g *GeminiController) RecommendContent(ctx *gin.Context) {
	sampleBlog := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	resp, err := g.model.RecommendContent(sampleBlog.Title, sampleBlog.Tags)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
}

// RecommendTags godoc
// @Summary Get tags recommendations
// @Description Get AI-generated tags recommendations based on title and content
// @Tags AI Recommendations
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for tags recommendation"
// @Success 200 {object} map[string][]string "response: Recommended tags"
// @Failure 406 {object} map[string]string "error: Invalid input data"
// @Failure 500 {object} map[string]string "error: AI model error"
// @Router /recommend/tags [post]
func (g *GeminiController) RecommendTags(ctx *gin.Context) {
	sampleBlog := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	resp, err := g.model.RecommendTags(sampleBlog.Title, sampleBlog.Content)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
}

// Summarize godoc
// @Summary Get blog summary
// @Description Get AI-generated summary of a blog based on title, content, and tags
// @Tags AI Recommendations
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for summary generation"
// @Success 200 {object} map[string]string "response: Blog summary"
// @Failure 406 {object} map[string]string "error: Invalid input data"
// @Failure 500 {object} map[string]string "error: AI model error"
// @Router /summarize [post]
func (g *GeminiController) Summarize(ctx *gin.Context) {
	sampleBlog := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	resp, err := g.model.Summarize(sampleBlog)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
}

// Refine godoc
// @Summary Refine blog content
// @Description Get AI-generated refined content based on input content
// @Tags AI Recommendations
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for content refinement"
// @Success 200 {object} map[string]string "response: Refined content"
// @Failure 406 {object} map[string]string "error: Invalid input data"
// @Failure 500 {object} map[string]string "error: AI model error"
// @Router /refine [post]
func (g *GeminiController) Refine(ctx *gin.Context) {
	sampleBlog := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	resp, err := g.model.Refine(sampleBlog.Content)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": resp})
}

// Validate godoc
// @Summary Validate blog content
// @Description Validate blog content using AI model
// @Tags AI Recommendations
// @Accept json
// @Produce json
// @Param input body infrastructure.Data true "Input data for validation"
// @Success 200 {object} map[string]string "response: Validation result"
// @Failure 406 {object} map[string]string "error: Invalid input data"
// @Failure 500 {object} map[string]string "error: AI model error"
// @Router /validate [post]
func (g *GeminiController) Validate(ctx *gin.Context) {
	sampleBlog := infrastructure.Data{}
	if err := ctx.ShouldBindJSON(&sampleBlog); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": "none acceptable data"})
		return
	}
	err := g.model.Validate(sampleBlog)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"response": "ok"})
}
