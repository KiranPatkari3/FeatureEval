package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Feature struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Operation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Variation struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Formula string   `json:"formula"`
	Params  []string `json:"params"`
}

type EvalRequest struct {
	Username    string            `json:"username"`
	VariationID int               `json:"variation_id"`
	Inputs      map[string]string `json:"inputs"`
}

type EvalResponse struct {
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

type APIResponse struct {
	Status       bool        `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	TotalRecords int         `json:"total_records,omitempty"`
	Page         int         `json:"page,omitempty"`
	Size         int         `json:"size,omitempty"`
}

func respondJSON(w http.ResponseWriter, status bool, message string, data interface{}, totalRecords, page, size int) {
	w.Header().Set("Content-Type", "application/json")
	response := APIResponse{
		Status:       status,
		Message:      message,
		Data:         data,
		TotalRecords: totalRecords,
		Page:         page,
		Size:         size,
	}
	json.NewEncoder(w).Encode(response)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Invalid request body",
		})
		return
	}

	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", creds.Username, creds.Password)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "User already exists or error occurred",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Signup successful",
		"data": map[string]string{
			"username": creds.Username,
		},
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Invalid request body",
		})
		return
	}

	row := db.QueryRow("SELECT id FROM users WHERE username=$1 AND password=$2", creds.Username, creds.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": "Invalid credentials",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Login successful",
		"data": map[string]interface{}{
			"user_id":  id,
			"username": creds.Username,
		},
	})
}

func FeaturesHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size

	rows, err := db.Query("SELECT id, name FROM features LIMIT $1 OFFSET $2", size, offset)
	if err != nil {
		respondJSON(w, false, "Database error", nil, 0, page, size)
		return
	}
	defer rows.Close()

	var features []Feature
	for rows.Next() {
		var f Feature
		rows.Scan(&f.ID, &f.Name)
		features = append(features, f)
	}

	row := db.QueryRow("SELECT COUNT(*) FROM features")
	var total int
	row.Scan(&total)

	respondJSON(w, true, "Features fetched", features, total, page, size)
}

func OperationsHandler(w http.ResponseWriter, r *http.Request) {
	featureID := r.URL.Query().Get("feature_id")
	rows, err := db.Query("SELECT id, name FROM operations WHERE f_id = $1", featureID)
	if err != nil {
		respondJSON(w, false, "Database error", nil, 0, 0, 0)
		return
	}
	defer rows.Close()

	var ops []Operation
	for rows.Next() {
		var o Operation
		rows.Scan(&o.ID, &o.Name)
		ops = append(ops, o)
	}
	respondJSON(w, true, "Operations fetched", ops, len(ops), 1, len(ops))
}

func VariationsHandler(w http.ResponseWriter, r *http.Request) {
	opID := r.URL.Query().Get("operation_id")
	rows, err := db.Query("SELECT id, name, formula, params FROM variations WHERE o_id = $1", opID)
	if err != nil {
		respondJSON(w, false, "Database error", nil, 0, 0, 0)
		return
	}
	defer rows.Close()

	var variations []Variation
	for rows.Next() {
		var v Variation
		var paramsJSON string
		rows.Scan(&v.ID, &v.Name, &v.Formula, &paramsJSON)

		var paramMap map[string]string
		json.Unmarshal([]byte(paramsJSON), &paramMap)
		for k := range paramMap {
			v.Params = append(v.Params, k)
		}
		variations = append(variations, v)
	}
	respondJSON(w, true, "Variations fetched", variations, len(variations), 1, len(variations))
}

func EvaluateHandler(w http.ResponseWriter, r *http.Request) {
	var req EvalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondJSON(w, false, "Invalid input", nil, 0, 0, 0)
		return
	}

	var formula string
	var paramsJSON string
	var variationName, programName string
	var opID int

	row := db.QueryRow(`
		SELECT v.formula, v.params, v.name, o.name, o.id
		FROM variations v
		JOIN operations o ON v.o_id = o.id
		WHERE v.id = $1`, req.VariationID)
	err = row.Scan(&formula, &paramsJSON, &variationName, &programName, &opID)
	if err != nil {
		respondJSON(w, false, "Variation not found", nil, 0, 0, 0)
		return
	}

	expr, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		respondJSON(w, false, "Invalid formula", nil, 0, 0, 0)
		return
	}

	parameters := make(map[string]interface{})
	for k, v := range req.Inputs {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			respondJSON(w, false, "Invalid number input", nil, 0, 0, 0)
			return
		}
		parameters[k] = val
	}

	result, err := expr.Evaluate(parameters)
	if err != nil {
		respondJSON(w, false, "Evaluation error", nil, 0, 0, 0)
		return
	}

	resultFloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", result), 64)
	inputsJSON, _ := json.Marshal(req.Inputs)

	_, err = db.Exec(`
		INSERT INTO user_logs (category, program, inputs, result, status, variation_id, logged_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		"Math/String", programName, string(inputsJSON), resultFloat, "Success", req.VariationID, time.Now())
	if err != nil {
		log.Println("Error logging:", err)
	}

	respondJSON(w, true, "Evaluation successful", EvalResponse{Status: "success", Result: resultFloat}, 1, 1, 1)
}
