package routers

import (
	"github.com/yesetoda/BlogMate/config"
	"github.com/yesetoda/BlogMate/delivery/controllers"
	"github.com/yesetoda/BlogMate/gemini"
	"github.com/yesetoda/BlogMate/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func AddAIRoutes(r *gin.Engine, config config.Config, prompts infrastructure.Prompts) {
	model, err := gemini.NewGeminiModel(config.Gemini.ApiKey, config.Gemini.Model, prompts)
	if err != nil {
		log.Fatal(err)
	}
	aiController := controllers.NewAIController(model)
	aiRouteGroup := r.Group("/ai")
	{
		aiRouteGroup.POST("/recommend", aiController.RecommendBlog)
		aiRouteGroup.POST("/recommendTitle", aiController.RecommendTitle)
		aiRouteGroup.POST("/recommendContent", aiController.RecommendContent)
		aiRouteGroup.POST("/recommendTags", aiController.RecommendTags)
		aiRouteGroup.POST("/summarize", aiController.Summarize)
		aiRouteGroup.POST("/chat", aiController.Chat)
	}
}
