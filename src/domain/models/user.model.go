package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;auto_increament:true;index"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

