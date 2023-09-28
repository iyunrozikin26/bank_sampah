package config

import (
	"database/sql"
	"log"

	"github.com/iyunrozikin26/bank_sampah.git/src/settings"
)

func DbSqlConnect() (*sql.DB, error) {
	database := settings.GoDotEnvVariable("DBNAMESERVER")
	host := settings.GoDotEnvVariable("DBHOST")
	port := settings.GoDotEnvVariable("DBPORT")
	dbname := settings.GoDotEnvVariable("DBNAME")
	username := settings.GoDotEnvVariable("DBUSERNAME")
	password := settings.GoDotEnvVariable("DBPASSWORD")

	db, err := sql.Open(database, username+":"+password+"@tcp("+host+":"+port+")"+"/"+dbname)

	if err != nil {
		return nil, err
	}
	log.Println("=== Koneksi database/sql berhasil ===")
	return db, nil
}
