package utils

import (
	"log"
	"os"
	"time"
)

func WriteLog(msg string) {
	now := time.Now()
	fileName := now.Format("2006.01.02-15.04") + "log.txt"

	// Open the log file in append mode, create if it doesn't exist
	file, err := os.OpenFile("./"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		// If file doesn't exist, create it
		if os.IsNotExist(err) {
			file, err = os.Create("log.txt")
			if err != nil {
				log.Fatalf("Failed to create log file: %v", err)
			}
		} else {
			log.Fatalf("Failed to open log file: %v", err)
		}
	}
	defer file.Close()

	// Write the log message to the file
	if _, err := file.WriteString(msg + "\n"); err != nil {
		log.Fatalf("Failed to write to log file: %v", err)
	}
}
