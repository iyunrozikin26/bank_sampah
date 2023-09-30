package user

import "time"

type UpdateUserDTO struct {
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)" validate:"email"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Role      string    `json:"role" gorm:"type:varchar(25)"`
	UpdatedAt time.Time `json:"updated_at"`
}
