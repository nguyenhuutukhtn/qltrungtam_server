package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/")
	{
		client.POST("/login", controllers.HandleLogin)
		client.POST("/create_account", controllers.HandleCreateAccount)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8001") // Ứng dụng chạy tại cổng 8081
}
