package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/ultimateGym/models"
)

type LoginServiceInterface interface {
	Login(login models.UserLoginDTO) (string, error)
}

type LoginController struct {
	userService LoginServiceInterface
}

func NewLoginController(u LoginServiceInterface) *LoginController {
	return &LoginController{userService: u}
}

func (l *LoginController) Login(c *gin.Context) {
	var login models.UserLoginDTO
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := l.userService.Login(login)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}