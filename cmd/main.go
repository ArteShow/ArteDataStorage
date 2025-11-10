package main

import "github.com/ArteShow/ArteDataStorage/pkg/logger"

func main() {
	logger.Init()
	logger.Log.Println("Initializing logger")
}
