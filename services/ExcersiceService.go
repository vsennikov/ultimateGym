package services

import "github.com/vsennikov/ultimateGym/models"

type ExerciseDBInterface interface {
	CreateExercise(exercise *models.Exercise) (uint, error)
	DeleteExercise(id uint) error
	UpdatedExercise(id uint, updates map[string]interface{}) error
	GetExerciseByID(id uint) (*models.Exercise, error)
	GetExerciseByName(name string) (*models.Exercise, error)
	GetAllExercises() ([]models.Exercise, error)
	GetAllUserExercises(userID uint) ([]models.Exercise, error)
	GetExercisesByType(muscle_group string) ([]models.Exercise, error)
}

type ExerciseService struct {
	repository ExerciseDBInterface
}

func NewExerciseService(r ExerciseDBInterface) *ExerciseService {
	return &ExerciseService{repository: r}
}


func (e *ExerciseService) CreateExercise(exercise models.ExerciseDTO) (uint, error) {
	newExercise := &models.Exercise{
		Name: exercise.Name,
		MuscleGroup: exercise.MuscleGroup,
		Description: exercise.Description,
		VideoURL: exercise.VideoURL,
		UserID: exercise.UserID,
	}
	return e.repository.CreateExercise(newExercise)
}