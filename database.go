package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Replace with your actual database name
const dataSource = "root:admin@321@tcp(127.0.0.1:3306)/logs"

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", dataSource)
}

// Optional: Log to DB for debugging or audit trails
func writeDBLog(category, program, inputs, result, status string) {
	db, err := connectDB()
	if err != nil {
		fmt.Println("DB connection failed:", err)
		return
	}
	defer db.Close()

	query := `INSERT INTO user_logs (category, program, inputs, result, status) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, category, program, inputs, result, status)
	if err != nil {
		fmt.Println("Failed to insert log:", err)
	} else {
		fmt.Println("✅ Log stored in database.")
	}
}
