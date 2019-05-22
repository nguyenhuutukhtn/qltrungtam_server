package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"log"
)

func GetAllCourses(c *gin.Context) {

	type CourseResponse struct {
		Id            int    `json:"Id,omitempty"`
		Name          string `json:"Name,omitempty"`
		Type          string `json:"Type"`
		StartDay      string `json:"StartDay"`
		Teacher       string `json:"Teacher"`
		FinishDay     string `json:"FinishDay"`
		StudentNumber int    `json:"StudentNumber"`
		Fee           int    `json:"Fee"`
		DiscountInfo  string `json:"DiscountInfo"`
	}

	db := database.DBConn()
	rows, err := db.Query("SELECT COURSE.Id,COURSE.Name,COURSE.Type,COURSE.StartDay,STAFF.Name,COURSE.FinishDay,COURSE.StudentNumber,COURSE.Fee,COURSE.DiscountInfo FROM COURSE,STAFF WHERE COURSE.TeacherId=STAFF.Id")
	if err != nil {

		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	listAllCourses := []CourseResponse{}
	for rows.Next() {
		var c CourseResponse
		err = rows.Scan(&c.Id, &c.Name, &c.Type, &c.StartDay, &c.Teacher, &c.FinishDay, &c.StudentNumber, &c.Fee, &c.DiscountInfo)
		if err != nil {
			log.Fatal(err)
		}
		listAllCourses = append(listAllCourses, c)
	}

	c.JSON(200, listAllCourses)
	defer rows.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
