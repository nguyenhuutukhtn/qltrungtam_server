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

func AddNewStudent(c *gin.Context) {

	//Get user request, parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := model.StudentInfoRequest{}
	json.Unmarshal([]byte(jsonStringRequest), &userRequest)
	fmt.Println(jsonStringRequest)

	//save to database

	db := database.DBConn()
	stmtIns, insertErr := db.Query("INSERT INTO STUDENT(Id,Name,Birthday,Grade,Gender,PhoneNum,Email,School,ParentName,ParentPhoneNum) VALUES (0,'" + userRequest.Name + "','" + userRequest.Birthday + "','" + strconv.Itoa(userRequest.Grade) + "','" + userRequest.Gender + "','" + userRequest.PhoneNum + "','" + userRequest.Email + "','" + userRequest.School + "','" + userRequest.ParentName + "','" + userRequest.ParentPhoneNum + "');")
	if insertErr != nil {
		log.Fatal(insertErr)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	c.JSON(200, userRequest)
	defer stmtIns.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
