package main

import (
	"fmt"
	"github.com/luankkobs/goapi/database"
	"github.com/luankkobs/goapi/models"
	"github.com/luankkobs/goapi/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{ID: 1, Name: "Luan", History: "I'm a developer"},
		{ID: 2, Name: "John", History: "I'm a QA"},
	}
	database.DatabaseConnection()
	fmt.Println("Starting the application...")
	routes.HandleRequests()
}
