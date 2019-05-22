package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/model"
	"log"
)

func GetAllStudent(c *gin.Context) {

	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM STUDENT")
	if err != nil {

		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	listAllStudent := []model.Student{}
	for rows.Next() {
		var s model.Student
		err = rows.Scan(&s.Id, &s.Name, &s.Birthday, &s.Grade, &s.Gender, &s.PhoneNum, &s.Email, &s.School, &s.ParentName, &s.ParentPhoneNum)
		if err != nil {
			log.Fatal(err)
		}
		listAllStudent = append(listAllStudent, s)
	}

	c.JSON(200, listAllStudent)
	defer rows.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
