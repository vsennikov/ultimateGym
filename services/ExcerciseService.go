package services

import (
	"errors"

	"github.com/vsennikov/ultimateGym/models"
)

type ExerciseDBInterface interface {
	CreateExercise(exercise *models.Exercise) (uint, error)
	DeleteExercise(id uint) error
	UpdatedExercise(id uint, updates map[string]interface{}) error
	GetExerciseByID(id uint) (*models.Exercise, error)
	GetExerciseByName(name string, userID uint) (*models.Exercise, error)
	GetAllExercises(userID uint) ([]models.Exercise, error)
	GetAllUserExercises(userID uint) ([]models.Exercise, error)
	GetExercisesByType(userId uint, muscle_group string) ([]models.Exercise, error)
	GetUserExercisesByType(userId uint, muscle_group string) ([]models.Exercise, error)
	GetUserExerciseByName (name string, userID uint) (*models.Exercise, error)
}

type ExerciseService struct {
	repository ExerciseDBInterface
}

func NewExerciseService(r ExerciseDBInterface) *ExerciseService {
	return &ExerciseService{repository: r}
}


func (e *ExerciseService) CreateExercise(exercise models.ExerciseDTO, userId uint) (uint, error) {
	newExercise := &models.Exercise{
		Name: exercise.Name,
		MuscleGroup: exercise.MuscleGroup,
		Description: exercise.Description,
		VideoURL: exercise.VideoURL,
		UserID: userId,
	}
	_, err := e.repository.GetExerciseByName(exercise.Name, userId)
	if err != nil {
	return e.repository.CreateExercise(newExercise)
	}
	return 0, errors.New("exercise already exists")
}

func (e *ExerciseService) GetAllExercises(userId uint) ([]models.ExerciseDTO, error) {
	
	var exercisesDTO []models.ExerciseDTO
	var exercises []models.Exercise
	exercises, err := e.repository.GetAllExercises(userId)
	if err != nil {
		return nil, err
	}
	for _, exercise := range exercises {
		exercisesDTO = append(exercisesDTO, e.transferExercise(exercise))
	}
	return exercisesDTO, nil
}

func (e *ExerciseService) GetAlluserExercises(userId uint) ([]models.ExerciseDTO, error) {
	var exercisesDTO []models.ExerciseDTO
	var exercises []models.Exercise
	exercises, err := e.repository.GetAllUserExercises(userId)
	if err != nil {
		return nil, err
	}
	for _, exercise := range exercises {
		exercisesDTO = append(exercisesDTO, e.transferExercise(exercise))
	}
	return exercisesDTO, nil
}

func (e *ExerciseService) GetAllExercisesByType(userId uint, muscle_group string) ([]models.ExerciseDTO, error) {
	var exercisesDTO []models.ExerciseDTO
	var exercises []models.Exercise
	exercises, err := e.repository.GetExercisesByType(userId, muscle_group)
	if err != nil {
		return nil, err
	}
	for _, exercise := range exercises {
		exercisesDTO = append(exercisesDTO, e.transferExercise(exercise))
	}
	return exercisesDTO, nil
}

func (e *ExerciseService) GetAllUserExercisesByType(userId uint, muscle_group string) ([]models.ExerciseDTO, error) {
	var exercisesDTO []models.ExerciseDTO
	var exercises []models.Exercise
	exercises, err := e.repository.GetUserExercisesByType(userId, muscle_group)
	if err != nil {
		return nil, err
	}
	for _, exercise := range exercises {
		exercisesDTO = append(exercisesDTO, e.transferExercise(exercise))
	}
	return exercisesDTO, nil
}

func (e *ExerciseService) GetExerciseByName(name string, userId uint) (models.ExerciseDTO, error) {
	exercise, err := e.repository.GetExerciseByName(name, userId)
	if err != nil {
		return models.ExerciseDTO{}, err
	}
	return e.transferExercise(*exercise), nil
}

func (e *ExerciseService) UpdatedExercise(name string, userId uint, updates map[string]interface{}) error {
	exercise, err := e.repository.GetUserExerciseByName(name, userId)
	if err != nil {
		return err
	}
	err = e.repository.UpdatedExercise(exercise.ID, updates)
	if err != nil {
		return nil
	}
	return nil
}

func (e *ExerciseService) DeleteExercise(name string, userId uint) error {
	exercise, err := e.repository.GetUserExerciseByName(name, userId)
	if err != nil {
		return err
	}
	return e.repository.DeleteExercise(exercise.ID)
}

func (e *ExerciseService) transferExercise(exercise models.Exercise) models.ExerciseDTO {
	return models.ExerciseDTO{
		ID: exercise.ID,
		Name: exercise.Name,
		MuscleGroup: exercise.MuscleGroup,
		Description: exercise.Description,
		VideoURL: exercise.VideoURL,
	}
}