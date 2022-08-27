package main

import (
	"crowdfunding/auth"
	"crowdfunding/handler"
	"crowdfunding/helper"
	"crowdfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	
	authService := auth.NewService()


	userHandler := handler.NewUserHandler(userService, authService)
	
	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

	}


// karena ingin function validation token & get user by id maka harus begini bentuk functionnya 
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") { //apakah di authheader ada kata bearer
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// default Bearer tookentokentoken karena kita ingin ambil token jadi harus di splitt
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1] //[Bearer, tokentokentoken]
		}

		// validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

}

}

	// INI ADALAH FUNCTION LAMA DAN YANG TERBARU ADA LAH DIATAS

	// func authMiddleware(c *gin.Context) {

	// 	authHeader := c.GetHeader("Authorization")

	// 	if !strings.Contains(authHeader, "Bearer") { //apakah di authheader ada kawa bearer
	// 		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized,  "error", nil)
	// 		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
	// 		return
	// 	}

	// 	// default Bearer tookentokentoken karena kita ingin ambil token jadi harus di splitt
	// 	tokenString := ""
	// 	arrayToken := strings.Split(authHeader, " ")
	// 	if len(tokenString) == 2 {
	// 		tokenString = arrayToken[1] //[Bearer, tokentokentoken]
	// 	}

	// 	// validasi token
	// 	token, err := 


	// LANGKAH LANGKAH MIDDLEWARE MENGGUNAKAN JWT

	// ambil nilai header Authorization: Bearer tokentoken/isi dari generate token
	// dari header authorization, ambil nilai dari tokennya saja
	// validasi token
	// ambil user_id
	// ambil user dari db berdasarkan user_id lewat service
	// kita set context isinya user


	// =============================
	// TEST JWT TOKEN VALIDATIOn
	// =============================


	// token,err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.zMk3KISsMiv_cBrc5H2oxyT0JXeGJUwPm4VDY0C-yXc")

	// if err != nil {
	// 	fmt.Println("Error JWT Valdiation")
	// }

	// if token.Valid {
	// 	fmt.Println("Successsssssssssss valid")
	// } else {
	// 	fmt.Println("Invalidddddddddddddddd")
	// }



	// =============================
	// TEST JWT TOKEN
	// =============================

	// fmt.Println(authService.GenerateToken(1001))


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


	

	

	


	