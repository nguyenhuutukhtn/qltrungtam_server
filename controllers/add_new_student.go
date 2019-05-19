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
	stmtIns, insert_err := db.Query("INSERT INTO HOCVIEN(id,HoTen,NgaySinh,Lop,GioiTinh,SDT,Email,Truong,HoTenPhuHuynh,SDTPhuHuynh) VALUES (0,'" + userRequest.HoTen + "','" + userRequest.NgaySinh + "','" + strconv.Itoa(userRequest.Lop) + "','" + userRequest.GioiTinh + "','" + userRequest.SDT + "','" + userRequest.Email + "','" + userRequest.Truong + "','" + userRequest.HoTenPhuHuynh + "','" + userRequest.SDTPhuHuynh + "');")
	if insert_err != nil {
		//log.SetOutput(gin.DefaultWriter) // You may need this
		//log.Println("test")
		//fmt.Println("avcss")
		log.Fatal(insert_err)
		c.JSON(500, gin.H{
			"messages": "not success",
		})
	}

	c.JSON(200, userRequest)
	defer stmtIns.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
