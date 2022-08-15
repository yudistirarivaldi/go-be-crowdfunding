package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

		for _, e := range err.(validator.ValidationErrors) { //ngubah err ke validator.ValidationErrors
			errors = append(errors, e.Error())
		}

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