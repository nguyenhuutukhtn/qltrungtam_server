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
		client.GET("/get_all_student", controllers.GetAllStudent)
		client.PUT("/update_student", controllers.UpdateStudent)
		client.GET("/get_all_courses", controllers.GetAllCourses)
		client.GET("/get_all_teachers", controllers.GetAllTeachers)
		client.POST("/add_new_course", controllers.AddNewCourse)
		client.GET("/get_opening_courses", controllers.GetOpeningCourses)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8001") // Ứng dụng chạy tại cổng 8001
}
