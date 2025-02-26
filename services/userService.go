package services

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vsennikov/ultimateGym/models"
)



type UserDBInterface interface {
	CreateUser(user *models.User) (uint, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserService struct {
	repository UserDBInterface
}

func NewUserService(r UserDBInterface) *UserService {
	return &UserService{repository: r}
}

func (u *UserService) CreateUser(user models.UserDTO) (uint, error) {
	if _, err := u.repository.GetUserByEmail(user.Email); err == nil {
		return 0, errors.New("user already exists")
	}
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

func (u *UserService) DeleteUser(id uint) error {
	return u.repository.DeleteUser(id)
}

func (u *UserService) Login(login models.UserLoginDTO) (string, error) {

	user, err := u.repository.GetUserByEmail(login.Email)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !checkPasswordHash(login.Password, user.PasswordHash) {
		return "", errors.New("invalid password")
	}
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

func (u *UserService) DecodeToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}
	return 0, errors.New("invalid token")
}