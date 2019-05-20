package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/model"
	"log"
)

func GetAllStudent(c *gin.Context) {

	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM HOCVIEN")
	if err != nil {
		//log.SetOutput(gin.DefaultWriter) // You may need this
		//log.Println("test")
		//fmt.Println("avcss")
		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	listAllStudent := []model.Student{}
	for rows.Next() {
		var s model.Student
		err = rows.Scan(&s.Id, &s.HoTen, &s.NgaySinh, &s.Lop, &s.GioiTinh, &s.SDT, &s.Email, &s.Truong, &s.HoTenPhuHuynh, &s.SDTPhuHuynh)
		if err != nil {
			log.Fatal(err)
		}
		listAllStudent = append(listAllStudent, s)
	}

	c.JSON(200, listAllStudent)
	defer rows.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
