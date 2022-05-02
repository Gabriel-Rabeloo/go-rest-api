package main

import (
	"fmt"

	"github.com/Gabriel-Rabeloo/go-rest-api/database"
	"github.com/Gabriel-Rabeloo/go-rest-api/models"
	"github.com/Gabriel-Rabeloo/go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}

	database.ConnectDataBase()
	fmt.Println("Starting server Rest with go")
	routes.HandleRequest()
}
