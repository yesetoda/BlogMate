package main

import (
	"context"

	"github.com/yesetoda/BlogMate/config"
	"github.com/yesetoda/BlogMate/delivery/controllers"
	_ "github.com/yesetoda/BlogMate/delivery/docs"
	router "github.com/yesetoda/BlogMate/delivery/routers"
	"github.com/yesetoda/BlogMate/infrastructure"
	"github.com/yesetoda/BlogMate/repository"
	usecase "github.com/yesetoda/BlogMate/usecases"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	config_mongo, err := config.LoadConfig()
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	clientOptions := options.Client().ApplyURI(config_mongo.Database.Uri)
	if err != nil {
		panic(err)
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	blogCollections := client.Database("Blog-Mate").Collection("Blogs")
	userCollections := client.Database("Blog-Mate").Collection("Users")
	commentCollections := client.Database("Blog-Mate").Collection("Comments")
	replyCollections := client.Database("Blog-Mate").Collection("Replies")
	// _ = client.Database("BlogAPI").Collection("Tokens")
	blogRepo := repository.NewBlogRepository(mongoifc.WrapClient(client), mongoifc.WrapCollection(blogCollections), mongoifc.WrapCollection(commentCollections), mongoifc.WrapCollection(replyCollections))
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(*blogUsecase)
	authController := infrastructure.NewAuthController(blogRepo)
	userRepo := repository.NewUserRepository(mongoifc.WrapCollection(userCollections))
	userUsecase, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		panic(err)
	}
	UserController := controllers.NewUserController(userUsecase)
	prompts, err := infrastructure.LoadPrompt("prompts.json")
	if err != nil {
		panic(err)
	}
	Router := router.NewMainRouter(*UserController, *blogController, authController,*config_mongo,prompts)
	Router.GinBlogRouter()
}
