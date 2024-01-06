package entity

import (
	"time"
)

type User struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Age int `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}