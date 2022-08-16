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
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()























	// =============================
	// TEST UPLOAD AVATAR IN SERVICE
	// =============================

	// userService.SaveAvatar(6, "images/1-profile.png")
	
	// =================================================
	// CEK EMAIL TERSEDIA ATAU TIDAK MENGGUNAKAN SERVICE
	// =================================================

	// input := user.CheckEmailInput {
	// 	Email: "pesulapmerah123@gmail.com",
	// }

	// bool, err := userService.IsEmailAvailable(input)
	// if err != nil {
	// 	fmt.Println("Gagal")
	// }
	
	// fmt.Println(bool)

	// ====================================
	// TEST NYARI EMAIL MENGGUNAKAN SERVICE
	// ====================================

	// input := user.LoginInput {
	// 	Email: "yudistira@gmail.com",
	// 	Password: "yudistirar626",
	// }

	// user, err := userService.LoginUser(input)

	// if err != nil {
	// 	fmt.Println("Gagal Login")
	// 	fmt.Println(err.Error())
	// return
	// }
	
	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// =========================================
	// TEST FIND BY EMAIL MENGGUNAKAN REPOSITORY
	// =========================================

	// userByEmail, err := userRepository.FindByEmail("samsudin@gmail.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(userByEmail.Name)

	// ====================================
	// TEST CREATE USER MENGGUNAKAN SERVICE
	// ====================================

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Pesulap merah"
	// userInput.Occupation = "Pesulap"
	// userInput.Email = "pesulapmerah@gmail.com"
	// userInput.Password = "12345"

	// userService.RegisterUser(userInput)

	// =======================================
	// TEST CREATE USER MENGGUNAKAN REPOSITORY
	// =======================================

	// user := user.User {
	// 	Name : "Gus Samsudin",
	// 	Occupation: "Padepokna Nur Dzat",
	// 	Email: "samsudin@gmail.com",
	
	// }

	// userRepository.Save(user)


	

	

	

}
