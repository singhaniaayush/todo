package main

import (
	"todo/server/routers"
	"todo/server/utils"

	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {

	// getting default router
	route := gin.Default()

	// intializing routes
	routers.Init(route)

	err := utils.ConnectDataBase()

	if err != nil {
		panic(err)
	}

	// setting up gin server instance
	server := &http.Server{
		Addr:    ":5000",
		Handler: route,
	}
	// starting server
	server.ListenAndServe()
}
