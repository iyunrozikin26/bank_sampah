package config

import (
	"log"

	"github.com/iyunrozikin26/bank_sampah.git/src/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBGormConnect() *gorm.DB {
	host := settings.GoDotEnvVariable("DBHOST")
	port := settings.GoDotEnvVariable("DBPORT")
	dbname := settings.GoDotEnvVariable("DBNAME")
	username := settings.GoDotEnvVariable("DBUSERNAME")
	password := settings.GoDotEnvVariable("DBPASSWORD")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("=== Tidak dapat terkoneksi ke database ===")
	}
	log.Println("=== Koneksi database/gorm berhasil ===")
	return db
}
