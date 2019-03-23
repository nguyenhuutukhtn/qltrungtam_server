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

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func HandleLogin(c *gin.Context) {
	//get user request and parse to model
	jsonStringRequest, _ := ioutil.ReadAll(c.Request.Body)
	userRequest := UserLoginRequest{}
	json.Unmarshal([]byte(jsonStringRequest), &userRequest)

	//get hash password from database
	db := database.DBConn()
	query_str := "SELECT password FROM user WHERE username = '" + userRequest.Username + "'"
	fmt.Println(query_str)
	rows, err := db.Query("SELECT password FROM user WHERE username = '" + userRequest.Username + "'")
	if err != nil {
		log.Fatal(err)
		log.SetOutput(gin.DefaultWriter) // You may need this
		log.Println("test")
		fmt.Println("avcss")
		c.JSON(500, gin.H{
			"messages": "username do not exist",
		})
	}

	for rows.Next() {
		var password string

		err = rows.Scan(&password)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
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
	c.JSON(200, userRequest)

	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
