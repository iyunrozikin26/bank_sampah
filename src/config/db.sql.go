package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iyunrozikin26/bank_sampah.git/src/settings"
)

func DBSqlConnect() (*sql.DB, error) {
	database := settings.GoDotEnvVariable("DBNAMESERVER")
	host := settings.GoDotEnvVariable("DBHOST")
	port := settings.GoDotEnvVariable("DBPORT")
	dbname := settings.GoDotEnvVariable("DBNAME")
	username := settings.GoDotEnvVariable("DBUSERNAME")
	password := settings.GoDotEnvVariable("DBPASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := sql.Open(database, dsn)

	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}
	log.Println("=== Koneksi database/sql berhasil ===")
	return db, nil
}
