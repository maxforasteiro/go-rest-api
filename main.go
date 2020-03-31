package main

import (
	"fmt"
	"./config"
	"./models"
	"./routes"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {
	config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run()
}
