package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/database"
	"github.com/nguyenhuutukhtn/anhngudongdo_server/model"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
)

func HandleCreateAccount(c *gin.Context) {

	//Get user request, parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := model.UserCreateAccountRequest{}
	json.Unmarshal([]byte(jsonStringRequest), &userRequest)
	fmt.Println(jsonStringRequest)

	//encrypt password
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), cost)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "password not allow",
		})
		return
	}

	//save to database if parse success

	db := database.DBConn()
	stmtIns, insert_err := db.Query("INSERT INTO user(id,username,name,password,role) VALUES (0,'" + userRequest.Username + "','" + userRequest.Name + "','" + string(hash) + "','" + userRequest.Role + "');")
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
