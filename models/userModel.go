package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username		string	`gorm:"column:username;unique;default:null"`
    Email			string	`gorm:"type:varchar(100);unique;default:null"`
    PasswordHash	string	`gorm:"type:text;default:null"`
    TelegramChatID	int64	`gorm:"uniqueIndex;default:null"`
    IsActive		bool	`gorm:"default:true"`
}

type UserDTO struct {
	Username		string	`json:"username"`
	Email			string	`json:"email"`
	Password		string	`json:"password"`
	TelegramChatID	int64	`json:"telegram_chat_id"`
}
