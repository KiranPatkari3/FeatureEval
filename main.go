package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Read DB connection string from env variable
	connStr := os.Getenv("DB_CONN")
	if connStr == "" {
		log.Fatal("❌ Environment variable DB_CONN is not set")
	}

	//  Connect to PostgreSQL
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to open DB connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	log.Println("✅ Connected to database.")

	// ✅ HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/api/signup", SignupHandler)
	mux.HandleFunc("/api/login", LoginHandler)
	mux.HandleFunc("/api/features", FeaturesHandler)
	mux.HandleFunc("/api/operations", OperationsHandler)
	mux.HandleFunc("/api/variations", VariationsHandler)
	mux.HandleFunc("/api/evaluate", EvaluateHandler)

	// ✅ Enable CORS
	handler := cors.AllowAll().Handler(mux)

	// ✅ Start the server
	fmt.Println("🚀 Server running at http://localhost:8081")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", handler))
}
