package infrastructure

import (
	"github.com/vsennikov/ultimateGym/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDB struct {
}

func getUserDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&models.User{})
	return db
}

func (u *UserDB) CreateUser(user *models.User) (uint, error) {
	err := getUserDB().Save(&user).Error
	return user.Model.ID, err
}

func (u *UserDB) DeleteUser(id uint) error {
	err := getUserDB().Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

func (u *UserDB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := getUserDB().Where("email = ?", email).First(&user).Error
	return &user, err
}
