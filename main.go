package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/controllers"
	"github.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {

	route := gin.Default()

	route.GET("/posts", controllers.PostIndex)
	route.POST("/posts", controllers.PostCreate)
	route.GET("/posts/:id", controllers.PostShow)

	route.Run()
}
