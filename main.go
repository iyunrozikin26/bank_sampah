package main

import (
	"github.com/iyunrozikin26/bank_sampah.git/src/config"
	routes "github.com/iyunrozikin26/bank_sampah.git/src/routes"
)

func main() {
	_, _ = config.DbSqlConnect()

	routes.Run()
}
