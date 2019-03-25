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

func HandleLogin(c *gin.Context) {
	//get user request and parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := model.UserLoginRequest{}
	json.Unmarshal([]byte(jsonStringRequest), &userRequest)

	//get hash password from database
	db := database.DBConn()
	query_str := "SELECT password FROM user WHERE username = '" + userRequest.Username + "'"
	fmt.Println(query_str)
	rows, err := db.Query("SELECT * FROM user WHERE username = '" + userRequest.Username + "'")
	if err != nil {
		log.Fatal(err)
		log.SetOutput(gin.DefaultWriter) // You may need this
		log.Println("test")
		fmt.Println("avcss")
		c.JSON(500, gin.H{
			"messages": "username do not exist",
		})
		db.Close()
		return
	}

	var loginResponse model.UserLoginResponse
	var password string
	for rows.Next() {

		err = rows.Scan(&loginResponse.Id, &loginResponse.Username, &loginResponse.Name, &password, &loginResponse.Role)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
			c.JSON(500, gin.H{
				"messages": "can not get data from database",
			})
			//log.Fatal(passwordErr)
			defer db.Close()
			return
		}

		fmt.Println(password)
		fmt.Println(userRequest.Password)

		passwordErr := bcrypt.CompareHashAndPassword([]byte(password), []byte(userRequest.Password))
		if passwordErr != nil {
			c.JSON(500, gin.H{
				"messages": "wrong user name or password",
			})
			//log.Fatal(passwordErr)
			panic(passwordErr.Error())
			defer db.Close()
			return
		}
	}

	rows, err = db.Query("SELECT id FROM user_role WHERE role_name = '" + loginResponse.Role + "'")
	if err != nil {
		log.Fatal(err)
		log.SetOutput(gin.DefaultWriter) // You may need this
		log.Println("test")
		fmt.Println("avcss")
		c.JSON(500, gin.H{
			"messages": "can not get role id from database",
		})
		db.Close()
		return
	}

	for rows.Next() {

		err = rows.Scan(&loginResponse.RoleId)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
			c.JSON(500, gin.H{
				"messages": "can not get data from database",
			})
			//log.Fatal(passwordErr)
			defer db.Close()
			return
		}
	}
	c.JSON(200, loginResponse)

	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
