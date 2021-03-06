package Seeders

import (
	"GoGin/Models"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

// Use go faker to seed data
func seedUsers(DB *gorm.DB) {
	DB.Create(&Models.User{
		UserName: gofakeit.Username(),
		Password: "123456",
		Deposit:  gofakeit.Float64(),
		Role:     "seller",
	})
}
