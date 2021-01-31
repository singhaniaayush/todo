package routers

import (
	"net/http"
	"todo/server/controller"

	"github.com/gin-gonic/gin"
)

// Init - Initiaze routes
func Init(route *gin.Engine) {

	route.GET("/", func(contect *gin.Context) {
		contect.JSON(http.StatusOK, gin.H{"data": "Welcome to our Todo Application"})
	})
	route.GET("/todo", controller.RetrieveAllTodos)
	route.GET("/todo/:id", controller.RetrieveTodoByID)
	route.DELETE("/todo/:id", controller.DeleteTodoByID)
	route.POST("/todo", controller.SaveTodo)
	route.PUT("/todo", controller.UpdateTodo)

}
