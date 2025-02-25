package models

import "gorm.io/gorm"

type Exercise struct {
    gorm.Model
    Name        string    `gorm:"type:varchar(255);not null" json:"name"`
    Description string    `gorm:"type:text" json:"description"`
    VideoURL    string    `gorm:"type:varchar(255)" json:"video_url"`
    MuscleGroup string    `gorm:"type:varchar(255)" json:"muscle_group"`
    UserID      uint      `gorm:"not null" json:"user_id"` 
}


type ExerciseDTO struct {
	Name        string	`json:"name" binding:"required"`
	Description string	`json:"description"`
	VideoURL    string	`json:"video_url"`
	MuscleGroup string	`json:"muscle_group" binding:"required"`
	UserID      uint	`json:"user_id""` //Точно мені потрібно?
}

type ExerciseTgDTO struct {
	BotID			int64	`json:"bot_id" binding:"required"`
	TelegramChatID	int64	`json:"telegram_chat_id" binding:"required"`
	Name        	string	`json:"name" binding:"required"`
	Description 	string	`json:"description"`
	VideoURL    	string	`json:"video_url"`
	MuscleGroup 	string	`json:"muscle_group" binding:"required"`
}