package services

import "time"

type UpdateProductDTO struct {
	Title     string    `json:"title" gorm:"type:varchar(255)" validate:"required"`
	Category  string    `json:"category" gorm:"type:varchar(255)"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	UpdatedAt time.Time `json:"updated_at"`
}
