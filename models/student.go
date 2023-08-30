package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	StudentID string         `json:"student_id"`
	Major     string         `json:"major"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseSuccessGetAll struct {
	Message string    `json:"message"`
	Data    []Student `json:"data"`
}

type ResponseSuccess struct {
	Message string  `json:"message"`
	Data    Student `json:"data"`
}

type RequestStudent struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	StudentID string `json:"student_id"`
	Major     string `json:"major"`
}
