package models

import "gorm.io/gorm"

type Exercise struct {
    gorm.Model
    Name        string    `gorm:"type:varchar(255);not null"`
    Description string    `gorm:"type:text"`
    VideoURL    string    `gorm:"type:varchar(255)"`
    MuscleGroup string    `gorm:"type:varchar(255)"`
    UserID      uint      `gorm:"default:null` 
}

type ExerciseDTO struct {
	ID		  	uint	`json:"id"`
	Name        string	`json:"name" binding:"required"`
	Description string	`json:"description"`
	VideoURL    string	`json:"video_url"`
	MuscleGroup string	`json:"muscle_group" binding:"required"`
}

type ExerciseTgDTO struct {
	BotID			int64	`json:"bot_id" binding:"required"`
	TelegramChatID	int64	`json:"telegram_chat_id" binding:"required"`
	Name        	string	`json:"name" binding:"required"`
	Description 	string	`json:"description"`
	VideoURL    	string	`json:"video_url"`
	MuscleGroup 	string	`json:"muscle_group" binding:"required"`
}