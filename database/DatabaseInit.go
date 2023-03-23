package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	if errEnv := godotenv.Load(".env"); errEnv != nil {
		panic("cannot load .env")
	}
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to databse")
	}
	fmt.Println("connected to database")
}
