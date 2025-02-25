package infrastructure

import (
	"github.com/vsennikov/ultimateGym/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ExerciseDB struct {

}

func NewExerciseDB() *ExerciseDB {
	return &ExerciseDB{}
}

func getExerciseDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&models.Exercise{})
	return db
}

func (e *ExerciseDB) CreateExercise(exercise *models.Exercise) (uint, error) {
	err := getExerciseDB().Save(&exercise).Error
	return exercise.Model.ID, err
}

func (e *ExerciseDB) DeleteExercise(id uint) error {
	err := getExerciseDB().Where("id = ?", id).Delete(&models.Exercise{}).Error
	return err
}


func (e *ExerciseDB) UpdatedExercise(id uint, updates map[string]interface{}) error {
	err := getExerciseDB().Where("id = ?", id).Updates(updates).Error
	return err
}

func (e *ExerciseDB) GetExerciseByID(id uint) (*models.Exercise, error) {
	var exercise models.Exercise
	err := getExerciseDB().Where("id = ?", id).First(&exercise).Error
	return &exercise, err
}

func (e *ExerciseDB) GetExerciseByName(name string) (*models.Exercise, error) {
	var exercise models.Exercise
	err := getExerciseDB().Where("name = ?", name).First(&exercise).Error
	return &exercise, err
}

func (e *ExerciseDB) GetAllExercises() ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := getExerciseDB().Find(&exercises).Error
	return exercises, err
}

func (e *ExerciseDB) GetAllUserExercises(userID uint) ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := getExerciseDB().Where("user_id = ?", userID).Find(&exercises).Error
	return exercises, err
}

func (e *ExerciseDB) GetExercisesByType(muscle_group string) ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := getExerciseDB().Where("muscle_group = ?", muscle_group).Find(&exercises).Error
	return exercises, err
}
