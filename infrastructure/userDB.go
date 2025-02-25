package infrastructure

import (
	"errors"

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
	if user.Email != ""{
		_, err := u.getUserByEmail(user.Email)
		if err == nil {
			return 0, errors.New("user with this email already exists")
		}
	} else {
		_, err := u.getUserByTGID(user.TelegramChatID)
		if err == nil {
			return 0, errors.New("user with this telegram id already exists")
		}
	}
	err := getUserDB().Save(&user).Error
	return user.Model.ID, err
}

func (u *UserDB) DeleteUser(id uint) error {
	err := getUserDB().Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

func (u *UserDB) getUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := getUserDB().Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserDB) getUserByTGID(chatID int64) (*models.User, error) {
	var user models.User
	err := getUserDB().Where("telegram_chat_id = ?", chatID).First(&user).Error
	return &user, err
}