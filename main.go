package main

import (
	"github.com/nasrunrozikinmm/bank_sampah.git/src/config"
	routes "github.com/nasrunrozikinmm/bank_sampah.git/src/routes"
)

func main() {
	_, _ = config.DBSqlConnect() // checking sql database connect using sql query
	routes.Run()
}
