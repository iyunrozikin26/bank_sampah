package product

import "time"

type CreateProductDTO struct {
	Title     string    `json:"title" gorm:"type:varchar(255)" validate:"required"`
	Category  string    `json:"category" gorm:"type:varchar(255)" validate:"required"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
