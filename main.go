package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/controllers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/")
	{
		client.POST("/login", controllers.HandleLogin)
		client.POST("/create_account", controllers.HandleCreateAccount)
		client.POST("/add_new_student", controllers.AddNewStudent)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8001") // Ứng dụng chạy tại cổng 8081
}
