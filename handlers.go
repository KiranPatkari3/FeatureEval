package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
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

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "User already exists or error occurred", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Signup successful")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	row := db.QueryRow("SELECT id FROM users WHERE username=? AND password=?", creds.Username, creds.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Login successful")
}

func FeaturesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM features")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var features []Feature
	for rows.Next() {
		var f Feature
		rows.Scan(&f.ID, &f.Name)
		features = append(features, f)
	}
	json.NewEncoder(w).Encode(features)
}

func OperationsHandler(w http.ResponseWriter, r *http.Request) {
	featureID := r.URL.Query().Get("feature_id")
	rows, err := db.Query("SELECT id, name FROM operations WHERE f_id = ?", featureID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ops []Operation
	for rows.Next() {
		var o Operation
		rows.Scan(&o.ID, &o.Name)
		ops = append(ops, o)
	}
	json.NewEncoder(w).Encode(ops)
}

func VariationsHandler(w http.ResponseWriter, r *http.Request) {
	opID := r.URL.Query().Get("operation_id")
	rows, err := db.Query("SELECT id, name, formula, params FROM variations WHERE o_id = ?", opID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
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
	json.NewEncoder(w).Encode(variations)
}

func EvaluateHandler(w http.ResponseWriter, r *http.Request) {
	var req EvalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var formula string
	var paramsJSON string
	var variationName, programName string
	var opID int

	row := db.QueryRow("SELECT v.formula, v.params, v.name, o.name, o.id FROM variations v JOIN operations o ON v.o_id = o.id WHERE v.id = ?", req.VariationID)
	err = row.Scan(&formula, &paramsJSON, &variationName, &programName, &opID)
	if err != nil {
		http.Error(w, "Variation not found", http.StatusBadRequest)
		return
	}

	expr, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		http.Error(w, "Invalid formula", http.StatusInternalServerError)
		return
	}

	parameters := make(map[string]interface{})
	for k, v := range req.Inputs {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			http.Error(w, "Invalid number input", http.StatusBadRequest)
			return
		}
		parameters[k] = val
	}

	result, err := expr.Evaluate(parameters)
	if err != nil {
		http.Error(w, "Evaluation error", http.StatusInternalServerError)
		return
	}

	resultFloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", result), 64)

	inputsJSON, _ := json.Marshal(req.Inputs)
	_, err = db.Exec(`INSERT INTO user_logs (category, program, inputs, result, status, variation_id, logged_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		"Math/String", programName, string(inputsJSON), resultFloat, "Success", req.VariationID, time.Now())
	if err != nil {
		log.Println("Error logging:", err)
	}

	json.NewEncoder(w).Encode(EvalResponse{Status: "success", Result: resultFloat})
}
