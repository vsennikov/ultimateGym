package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsennikov/ultimateGym/models"
)

type ExerciseServiceInterface interface {
	CreateExercise(exercise models.ExerciseDTO) (uint, error)
}

type ExerciseController struct {
	exerciseService ExerciseServiceInterface
	userService UserServiceInterface

}

func NewExerciseController(s ExerciseServiceInterface, u UserServiceInterface) *ExerciseController {
	return &ExerciseController{exerciseService: s,
	userService: u}
}

func (e *ExerciseController) CreateExercise(c *gin.Context) {
	var req models.ExerciseDTO
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.UserID = userId
	exerciseID, err := e.exerciseService.CreateExercise(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercise_id": exerciseID})
}

func (e *ExerciseController) checkToken(token string) (uint, error) {
	if token == "" {
		return 0, errors.New("empty token")
	}
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	userId, err := e.userService.DecodeToken(token)
	if err != nil {
		return 0, errors.New("Unauthorized")
	}
	return userId, nil
}