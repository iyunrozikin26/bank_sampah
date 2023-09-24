package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	host := "127.0.0.1"
	port := "3306"
	dbname := "bank_sampah"
	username := "root"
	password := "root"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak dapat terkoneksi ke database")
	}
	fmt.Println("Koneksi database berhasil")
	return db
}
