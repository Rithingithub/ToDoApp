package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/Rithingithub/ToDoApp/db"
	"github.com/Rithingithub/ToDoApp/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errormessage := user.Validate()
	if len(errormessage) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
		return
	}
	collection := db.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Get user by ID
// @Description Get the user by ID
// @Tags Users
// @ID get-user-by-id
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	collection := db.GetDB().Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	c.JSON(http.StatusOK, user)
}
