package migration

import (
	"fmt"
	"log"
	"todoApps/database"
	"todoApps/models"
)

func RunMigrations() {
	err := database.DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Println("err")
	} else {
		fmt.Println("database migration success")
	}
}
