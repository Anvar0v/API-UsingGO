package main

import (
	"fmt"
	"movies_api/config"
	"movies_api/models"
)

func init() {
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.Movie{})
	fmt.Println("? Migration completed")
}
