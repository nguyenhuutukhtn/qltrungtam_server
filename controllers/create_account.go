package controllers

import (
	"../database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
)

type UserCreateAccountRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
	Role     string `json:""role,omitempty`
}

func HandleCreateAccount(c *gin.Context) {

	//Get user request, parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := UserCreateAccountRequest{}
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
