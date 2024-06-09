package main

import (
	"rental-ads-server/conf"
	"rental-ads-server/model"
	"rental-ads-server/server"
)

func main() {
	conf.Init()
	model.DBInit("horace:001357@tcp(localhost:3306)/rental_ads?charset=utf8mb4&parseTime=True")

	r := server.NewRouter()
	r.Run(":8080")
}
