package main

import (
	"github.com/ArteShow/ArteDataStorage/internal/database"
	"github.com/ArteShow/ArteDataStorage/pkg/logger"
)

func main() {
	logger.Init()
	logger.Log.Println("Initializing logger")

	logger.Log.Println("Preparing database")
	if err := database.Connect(); err != nil {
		logger.Log.Println("Failed to prepare the database")
	}
	database.Migrate()
	database.Close()
	logger.Log.Println("Database set up successful")

}
