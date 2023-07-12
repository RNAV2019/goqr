package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goqr struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	URL    string `json:"url" gorm:"not null"`
	QRCode string `json:"qrcode" gorm:"not null"`
}

func Setup() {
	// dsn := os.Getenv("DSN")
	dsn := "postgres://admin:test@" + os.Getenv("DB_HOST") + ":5432/admin?sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Goqr{})
	if err != nil {
		fmt.Println(err)
	}
}
