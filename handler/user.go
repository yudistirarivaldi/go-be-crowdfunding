package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	// tangkap input dari user melalui website
	// map input dari user ke struct RegisterUserInput
	// struct dia atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input) //mengubah struct ke json
	if err != nil {
		var errors []string

		errors = helper.FormatValidationError(err)

		errorMessage := gin.H{ "errors" : errors }

		response := helper.APIResponse("Register user gagal di tambahkan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //agar eksekusi stop di sini
	}

	 newUser, err := h.userService.RegisterUser(input)
	 
	 if err != nil {
		response := helper.APIResponse("Register user gagal di tambahkan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return //agar eksekusi stop di sini
	 }

	//  token, err := h.jwtService.GenerateToken(user)

	formatter := user.FormatUser(newUser, "tokenjwt")

	response := helper.APIResponse("Register User berhasil ditambahkan", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	// user memasukan input (email & password)
	// imnput di tangkap handler
	// mapping dari input user ke input struct
	// input struct di passing ke service
	// di service mencari dengan bantuan repository user dengan email tertentu
	// mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{ "errors" : errors}

		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()} //memanggil error yang ada di service
		response := helper.APIResponse("Login Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokenjwt")

	response := helper.APIResponse("Login Berhasil", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailability (c *gin.Context) {

	// ada input email dari user
	// input email di mapping ke struct input
	// struct input di passing ke service
	// service akan memanggil repository untuk ngecek apakah email sudah ada atau belum
	//  repository - db
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		response := helper.APIResponse("Check email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors" : errors}

		reponse := helper.APIResponse("Check email gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, reponse)
	}

	data := gin.H {
		"is_email_available" : isEmailAvailable,
	}

	metaMessage := "Email sudah Di Daftarkan"

	if isEmailAvailable { // if isEmailAvailable == true karena return defaultnya adalah false 
		metaMessage = "Email tersedia"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data) 
	c.JSON(http.StatusOK, response)

}