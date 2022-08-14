package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("Connection to Database Successful")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/user", userHandler.RegisterUser)

	router.Run()







	// test create user menggunakan service

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Pesulap merah"
	// userInput.Occupation = "Pesulap"
	// userInput.Email = "pesulapmerah@gmail.com"
	// userInput.Password = "12345"

	// userService.RegisterUser(userInput)


	// test Create User menggunakan repository

	// user := user.User {
	// 	Name : "Gus Samsudin",
	// 	Occupation: "Padepokna Nur Dzat",
	// 	Email: "samsudin@gmail.com",
	
	// }

	// userRepository.Save(user)


	

	

	

}
