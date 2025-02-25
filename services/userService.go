package services

import (
	"errors"

	"github.com/vsennikov/ultimateGym/models"
)



type UserDBInterface interface {
	CreateUser(user *models.User) (uint, error)
	DeleteUser(id uint) error
}

type UserService struct {
	repository UserDBInterface
}

func NewUserService(r UserDBInterface) *UserService {
	return &UserService{repository: r}
}

func (u *UserService) CreateUser(user models.UserDTO) (uint, error) {
	if user.TelegramChatID == 0 && user.Email == "" {
		return 0, errors.New("invalid user data")
	}
	if user.TelegramChatID != 0 && user.Email != "" {
		return 0, errors.New("invalid user data")
	}
	if user.TelegramChatID == 0 {
		valid := isEmailValid(user.Email)
		if !valid  {
			return 0, errors.New("invalid email")
		}
		if user.Password == "" {
			return 0, errors.New("invalid user data")
		}
		return u.CreateStandartUser(user)
	}
	return u.CreateTelegramUser(user)
}

func (u *UserService) CreateStandartUser(user models.UserDTO) (uint, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	newUser := &models.User{
		Username: user.Username,
		Email: user.Email,
		PasswordHash: hashedPassword,
		IsActive: true,
	}
	return u.repository.CreateUser(newUser)
}

func (u *UserService) CreateTelegramUser(user models.UserDTO) (uint, error) {
	if user.TelegramChatID == 0{
		return 0, errors.New("invalid user data")
	}
	newUser := &models.User{
		Username: user.Username,
		TelegramChatID: user.TelegramChatID,
		IsActive: true,
	}
	return u.repository.CreateUser(newUser)
}

func (u *UserService) DeleteUser(id uint) error {
	return u.repository.DeleteUser(id)
}
