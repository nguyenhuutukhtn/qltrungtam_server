package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/model"
	"io/ioutil"
	"log"
	"strconv"
)

func AddNewCourse(c *gin.Context) {

	type TimeTable struct {
		Course   model.Course `json:"course"`
		ListTime []string     `json:"listTime"`
	}

	//Get user request, parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := TimeTable{}
	json.Unmarshal([]byte(jsonStringRequest), &userRequest)
	fmt.Println(string([]byte(jsonStringRequest)))
	//fmt.Println(strconv.Itoa(len(userRequest.ListTime)) + "abc")

	//get teacher ID
	var teacherID int
	db := database.DBConn()
	rows, err := db.Query("SELECT STAFF.Id FROM STAFF WHERE STAFF.Name='" + userRequest.Course.Teacher + "';")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}
	for rows.Next() {
		err = rows.Scan(&teacherID)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Save course
	query, err := db.Query("INSERT INTO COURSE(Id,Name,Type,StartDay,TeacherId,FinishDay,StudentNumber,Fee,DiscountInfo) VALUES (0,'" + userRequest.Course.Name + "','" + userRequest.Course.Type + "','" + userRequest.Course.StartDay + "','" + strconv.Itoa(teacherID) + "','" + userRequest.Course.FinishDay + "','" + strconv.Itoa(userRequest.Course.StudentNumber) + "','" + strconv.Itoa(userRequest.Course.Fee) + "','" + userRequest.Course.DiscountInfo + "');")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	//get course id
	var courseID int
	rows, err = db.Query("SELECT MAX(Id) FROM COURSE;")
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}
	for rows.Next() {
		err = rows.Scan(&courseID)
		if err != nil {
			log.Fatal(err)
		}
	}

	//save time table
	for i := 0; i < len(userRequest.ListTime); i++ {
		query, err = db.Query("INSERT INTO TIMETABLE(Id,IdCourse,DateTime) VALUES (0,'" + strconv.Itoa(courseID) + "','" + userRequest.ListTime[i] + "');")
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{
				"messages": "not success",
			})
		}
	}

	c.JSON(200, userRequest)
	defer rows.Close()
	defer query.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
