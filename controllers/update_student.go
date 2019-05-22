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

func UpdateStudent(c *gin.Context) {

	//Get user request, parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := model.StudentInfoRequest{}
	json.Unmarshal([]byte(jsonStringRequest), &userRequest)
	fmt.Println(jsonStringRequest)

	//save to database

	db := database.DBConn()
	stmtUpd, updateErr := db.Query("UPDATE STUDENT SET Name='" + userRequest.Name + "',Birthday='" + userRequest.Birthday + "',Grade=" + strconv.Itoa(userRequest.Grade) + ",Gender='" + userRequest.Gender + "',PhoneNum='" + userRequest.PhoneNum + "',Email='" + userRequest.Email + "',School='" + userRequest.School + "',ParentName='" + userRequest.ParentName + "',ParentPhoneNum='" + userRequest.ParentPhoneNum + "' WHERE Id=" + strconv.Itoa(userRequest.Id))
	if updateErr != nil {
		//log.SetOutput(gin.DefaultWriter) // You may need this
		//log.Println("test")
		//fmt.Println("avcss")
		log.Fatal(updateErr)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	c.JSON(200, userRequest)
	defer stmtUpd.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
