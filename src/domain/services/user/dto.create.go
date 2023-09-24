package services

import "time"

type CreateUserDTO struct {
	Name      string    `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Email     string    `json:"email" gorm:"type:varchar(255)" validate:"email"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
