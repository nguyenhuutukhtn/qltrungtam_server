package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"log"
)

//GetAllCourses return
func GetAllTeachers(c *gin.Context) {

	type TeacherResponse struct {
		ID        int    `json:"Id,omitempty"`
		Name      string `json:"Name,omitempty"`
		Birthday  string `json:"Birthday"`
		IDCardNum string `json:"IdCardNum"`
		Gender    string `json:"Gender"`
		PhoneNum  string `json:"PhoneNum"`
		Email     string `json:"Email"`
		Address   string `json:"Address"`
		Role      string `json:"Role"`
	}

	db := database.DBConn()
	rows, err := db.Query("SELECT STAFF.Id,STAFF.Name,STAFF.Birthday,STAFF.IdCardNum,STAFF.Gender,STAFF.PhoneNum,STAFF.Email,STAFF.Address,STAFF.Role FROM STAFF WHERE STAFF.Role='Giáo viên'")
	if err != nil {

		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	listAllTeachers := []TeacherResponse{}
	for rows.Next() {
		var c TeacherResponse
		err = rows.Scan(&c.ID, &c.Name, &c.Birthday, &c.IDCardNum, &c.Gender, &c.PhoneNum, &c.Email, &c.Address, &c.Role)
		if err != nil {
			log.Fatal(err)
		}
		listAllTeachers = append(listAllTeachers, c)
	}

	c.JSON(200, listAllTeachers)
	defer rows.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
