package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/ultimateGym/models"
)

type UserServiceInterface interface {
	CreateUser(user models.UserDTO) (uint, error)
	DeleteUser(id uint) error
	DecodeToken(tokenString string) (uint, error)
}

type UserController struct {
	service UserServiceInterface
}

func NewUserController(s UserServiceInterface) *UserController {
	return &UserController{service: s}
}

func (r *UserController) Registration(c *gin.Context) {
	var req models.UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := r.service.CreateUser(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user_id": userID})
}
