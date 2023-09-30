package user

import (
	"gorm.io/gorm"
)

// kontrak methods
type UserRepository interface {
	FindAll() []User
	FindOne(id int) User
	Save(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) (*User, error)
	// SoftDelete(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (ur *UserRepositoryImpl) FindAll() []User {
	var users []User

	_ = ur.db.Find(&users) // menggunakan pointer agar merujuk ke memory yang sama
	return users

}
func (ur *UserRepositoryImpl) FindOne(id int) User {
	var user User
	_ = ur.db.First(&user, id)
	return user
}
func (ur *UserRepositoryImpl) Save(user User) (*User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (ur *UserRepositoryImpl) Update(user User) (*User, error) {
	result := ur.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (ur *UserRepositoryImpl) Delete(user User) (*User, error) {
	result := ur.db.Delete(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
