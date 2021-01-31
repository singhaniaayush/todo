package controller

import (
	"net/http"
	"todo/server/models"
	"todo/server/service"

	"github.com/gin-gonic/gin"
)

// RetrieveAllTodos - get all
func RetrieveAllTodos(context *gin.Context) {

	todos, err := service.RetrieveAllTodos()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"data": "No records found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": todos})
	return

}

// RetrieveTodoByID - retrieve todo by id
func RetrieveTodoByID(context *gin.Context) {
	todo, err := service.RetrieveTodoByID(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"data": "No records found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": todo})
	return
}

// DeleteTodoByID - delete todo by id
func DeleteTodoByID(context *gin.Context) {
	err := service.DeleteTodoByID(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"data": "No records found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "Deleted"})
	return
}

// SaveTodo - save todo
func SaveTodo(context *gin.Context) {
	var todo models.Todos

	err := context.BindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"data": "Error While binding request"})
		return
	}

	err = service.SaveTodo(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"data": "Error While saving"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": todo})
	return
}

// UpdateTodo - update todo
func UpdateTodo(context *gin.Context) {
	var todo models.Todos

	err := context.BindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"data": "Error While binding request"})
		return
	}

	err = service.UpdateTodo(&todo, context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"data": "Error While updating"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": todo})
	return
}
