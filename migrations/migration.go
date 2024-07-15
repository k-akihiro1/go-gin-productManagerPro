package main

import (
	"go-gin-productManagerPro/infra"
	"go-gin-productManagerPro/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		panic("Failed to migrate datebase")
	}
}
