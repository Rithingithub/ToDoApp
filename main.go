package main

import (
	"context"
	"log"
	"net/http"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Rithingithub/ToDoApp/controller"
	"github.com/Rithingithub/ToDoApp/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	client, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	r := gin.Default()
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Index route
	r.GET("/", func(ctx *gin.Context) {
		res := map[string]interface{}{
			"message": "Api is running",
			"ok":      true,
		}
		ctx.JSON(http.StatusOK, res)
	})

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", controller.CreateUser)
		userRoutes.GET("/:id", controller.GetUser)
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
