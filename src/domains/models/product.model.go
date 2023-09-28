package models

import (
	"time"
)

type Product struct {
	ID        int64     `json:"id" gorm:"primaryKey;auto_increament:true;index"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Category  string    `json:"category" gorm:"type:varchar(255)"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
