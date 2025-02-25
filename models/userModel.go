package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username		string	`gorm:"column:username;unique;default:null"`
    Email			string	`gorm:"type:varchar(100);unique;default:null"`
    PasswordHash	string	`gorm:"type:text;default:null"`
    TelegramChatID	int64	`gorm:"clumn:telegram_chat_id;unique;default:null"`
    IsActive		bool	`gorm:"default:true"`
}

type UserDTO struct {
	Username		string	`json:"username" binding:"required"`
	Email			string	`json:"email" binding:"required,email"`
	Password		string	`json:"password" binding:"required"`
}

type UserTgDTO struct {
	Username		string	`json:"username" binding:"required"`
	BotID			int64	`json:"bot_id" binding:"required"`
	TelegramChatID	int64	`json:"telegram_chat_id" binding:"required"`
}