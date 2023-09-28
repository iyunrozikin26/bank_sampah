package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// use before save password to db
func HashingAndSaltPassword(password string) string {
	hashing, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashing)
}

// use when user login and change password
func CompareAndCheckingPasswordHash(password, hashing string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashing), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
