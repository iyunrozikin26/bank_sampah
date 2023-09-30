package user

import (
	"time"

	"github.com/iyunrozikin26/bank_sampah.git/src/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;auto_increament:true;index"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Role      string    `json:"role" gorm:"type:varchar(25)"`
	Deleted   bool      `json:"deleted" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublicUser struct {
	ID        int64     `json:"id" gorm:"primaryKey;auto_increament:true;index"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Role      string    `json:"role" gorm:"type:varchar(25)"`
	Deleted   bool      `json:"deleted" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// Hook hashing password
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashingAndSaltPassword(u.Password)
	return
}

type Users []User

// So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:   u.ID,
		Name: u.Name,
		Role: u.Role,
	}
}
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}
