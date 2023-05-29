package main

import (
	"denny/gin_gorm_practise_dining/app/db"
	"denny/gin_gorm_practise_dining/app/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var d db.IDb

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "testing" {
		d = &db.Sqlite{}
	} else {
		d = &db.Mysql{}
	}

	db.InitDb(d.NewConnInfo())

	router := router.InitRouter()

	router.Run(":8081")
}
