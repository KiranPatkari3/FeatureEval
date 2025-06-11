package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	var err error

	dsn := "root:admin@321@tcp(127.0.0.1:3306)/logs?parseTime=true&timeout=5s"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB connection: %v", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	log.Println("✅ Connected to MySQL database.")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/signup", SignupHandler)
	mux.HandleFunc("/api/login", LoginHandler)
	mux.HandleFunc("/api/features", FeaturesHandler)
	mux.HandleFunc("/api/operations", OperationsHandler)
	mux.HandleFunc("/api/variations", VariationsHandler)
	mux.HandleFunc("/api/evaluate", EvaluateHandler)

	handler := cors.AllowAll().Handler(mux)

	fmt.Println("🚀 Server running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
