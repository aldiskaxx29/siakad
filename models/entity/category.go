package entity

import "time"

type Category struct {
	id uint `json:"id" gorm:"primrayKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}