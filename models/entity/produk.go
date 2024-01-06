package entity

import (
	"time"
)

type Produk struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Category string `json:"category"`
	Description string `json:"description"`
	CreatedAt time.Time
	UpdatedAt time.Time
}