package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/model"
	"log"
	"strconv"
)

//GetAllCourses return
func GetOpeningCourses(c *gin.Context) {

	type TimeTable struct {
		Course   model.Course `json:"course"`
		ListTime []string     `json:"listTime"`
	}

	response := []TimeTable{}

	db := database.DBConn()
	rows, err := db.Query("SELECT COURSE.Id,COURSE.Name,COURSE.Type,COURSE.StartDay,STAFF.Name as Teacher,COURSE.FinishDay,COURSE.StudentNumber,COURSE.Fee,COURSE.DiscountInfo from COURSE,STAFF where now() between STR_TO_DATE(StartDay,'%m/%d/%Y') and STR_TO_DATE(FinishDay,'%m/%d/%Y') and STAFF.Id=COURSE.TeacherId ; ")
	if err != nil {

		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	for rows.Next() {
		var c TimeTable
		err = rows.Scan(&c.Course.Id, &c.Course.Name, &c.Course.Type, &c.Course.StartDay, &c.Course.Teacher, &c.Course.FinishDay, &c.Course.StudentNumber, &c.Course.Fee, &c.Course.DiscountInfo)
		if err != nil {
			log.Fatal(err)
		}
		response = append(response, c)
	}

	for i := 0; i < len(response); i++ {
		rows, err = db.Query("select TIMETABLE.DateTime from TIMETABLE where TIMETABLE.IdCourse=" + strconv.Itoa(response[i].Course.Id) + ";")
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{
				"messages": "not success",
			})
		}

		list := []string{}
		for rows.Next() {
			var t string
			err = rows.Scan(&t)
			if err != nil {
				log.Fatal(err)
			}
			list = append(list, t)
		}
		response[i].ListTime = list
	}

	c.JSON(200, response)
	defer rows.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
