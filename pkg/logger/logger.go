package logger

import (
	"log"
	"os"
	"path/filepath"
)

var Log *log.Logger

func Init() {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get working directory: %v", err)
	}

	logsDir := filepath.Join(rootDir, "logs")
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Fatalf("cannot create logs directory: %v", err)
	}

	logFile := filepath.Join(logsDir, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("cannot open log file: %v", err)
	}

	Log = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}
