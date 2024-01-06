package entity

import (
	"time"
)

type Mahasiswa struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Study string `json:"study"`
	Photo string `json:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
}