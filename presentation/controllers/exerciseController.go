package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vsennikov/ultimateGym/models"
)

type ExerciseServiceInterface interface {
	CreateExercise(exercise models.ExerciseDTO, userId uint) (uint, error)
	GetAllExercises(userId uint) ([]models.ExerciseDTO, error)
	GetAlluserExercises(userId uint) ([]models.ExerciseDTO, error)
	GetAllExercisesByType(userId uint, muscle_group string) ([]models.ExerciseDTO, error)
	GetAllUserExercisesByType(userId uint, muscle_group string) ([]models.ExerciseDTO, error)
	GetExerciseByName(name string, userId uint) (models.ExerciseDTO, error)
	DeleteExercise(name string, userId uint) error
	UpdatedExercise(name string, userId uint, updates map[string]interface{}) error
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
	exerciseID, err := e.exerciseService.CreateExercise(req, userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercise_id": exerciseID})
}

func (e *ExerciseController) GetAllExercises(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	exercises, err := e.exerciseService.GetAllExercises(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercises": exercises})
}

func (e *ExerciseController) GetAllUserExercises(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	exercises, err := e.exerciseService.GetAlluserExercises(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercises": exercises})
}

func (e *ExerciseController) GetAllExercisesByType(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	muscle_group := c.Param("muscle_group")
	exercises, err := e.exerciseService.GetAllExercisesByType(userId, muscle_group)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercises": exercises})
}

func (e *ExerciseController) GetAllUserExercisesByType(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	muscle_group := c.Param("muscle_group")
	exercises, err := e.exerciseService.GetAllUserExercisesByType(userId, muscle_group)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercises": exercises})
}

func (e *ExerciseController) GetExerciseByName(c *gin.Context){
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	name := c.Param("name")
	exercise, err := e.exerciseService.GetExerciseByName(name, userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"exercise": exercise})
}

func (e *ExerciseController) UpdatedExercise(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	name := c.Param("name")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = e.exerciseService.UpdatedExercise(name, userId, updates)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "exercise updated"})
}

func (e *ExerciseController) DeleteExercise(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := e.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	name := c.Param("name")
	err = e.exerciseService.DeleteExercise(name, userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "exercise deleted"})
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