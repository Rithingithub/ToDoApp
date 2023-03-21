package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	//db := db()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		res := map[string]interface{}{
			"message": "Api is running",
			"ok":      true,
		}
		ctx.JSON(http.StatusOK, res)
	})

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func db() {
	panic("unimplemented")
}
