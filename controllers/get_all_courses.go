package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/model"
	"log"
)

//GetAllCourses return
func GetAllCourses(c *gin.Context) {

	db := database.DBConn()
	rows, err := db.Query("SELECT COURSE.Id,COURSE.Name,COURSE.Type,COURSE.StartDay,STAFF.Name,COURSE.FinishDay,COURSE.StudentNumber,COURSE.Fee,COURSE.DiscountInfo FROM COURSE,STAFF WHERE COURSE.TeacherId=STAFF.Id")
	if err != nil {

		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	listAllCourses := []model.Course{}
	for rows.Next() {
		var c model.Course
		err = rows.Scan(&c.Id, &c.Name, &c.Type, &c.StartDay, &c.Teacher, &c.FinishDay, &c.StudentNumber, &c.Fee, &c.DiscountInfo)
		if err != nil {
			log.Fatal(err)
		}
		listAllCourses = append(listAllCourses, c)
	}

	c.JSON(200, listAllCourses)
	defer rows.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
