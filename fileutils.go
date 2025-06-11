package main

import (
	"fmt"
	"os"
)

func writeLog(category, program, inputs, result, status string) {
	// Open file in append mode or create if doesn't exist
	f, err := os.OpenFile("user_activity_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	// Format your log entry (you can customize this)
	logLine := fmt.Sprintf("Category: %s | Program: %s | Inputs: %s | Result: %s | Status: %s\n",
		category, program, inputs, result, status)

	// Write the log entry to the file
	if _, err := f.WriteString(logLine); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}
