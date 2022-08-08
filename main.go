package main

import (
	"crowdfunding/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection to Database Successful")

	// var users []user.User
	// db.Find(&users)
	
	router := gin.Default()
	router.GET("/handler", handler)

	router.Run()
} 

func handler(c *gin.Context) {

	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to Database Successful")

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)


}