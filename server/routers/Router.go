package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init - Initiaze routes
func Init(route *gin.Engine) {

	route.GET("/", func(contect *gin.Context) {
		contect.JSON(http.StatusOK, gin.H{"data": "Welcome to our Todo Application"})
	})

}
