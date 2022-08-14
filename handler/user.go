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
		c.JSON(http.StatusBadRequest, err)
	}

	 user, err := h.userService.RegisterUser(input)
	 

	 if err != nil {
		c.JSON(http.StatusBadRequest, err)
	 }

	response := helper.APIResponse("User berhasil ditambahkan", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)


}