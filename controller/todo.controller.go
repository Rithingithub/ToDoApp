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

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errormessage := todo.Validate()
	if len(errormessage) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
		return
	}
	collection := db.GetDB().Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert todo list"})
		return
	}
	c.JSON(http.StatusOK, result)
}
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	collection := db.GetDB().Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todo list"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
