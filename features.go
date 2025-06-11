package main

/*
import (
	"encoding/json"
)

// // Feature maps to the features table
// type Feature struct {
// 	ID   int
// 	Name string
// }

// // Operation maps to the operations table
// type Operation struct {
// 	ID   int
// 	Name string
// 	FID  int
// }

// // Variation maps to the variations table
// type Variation struct {
// 	ID      int
// 	Name    string
// 	OID     int
// 	Formula string
// 	Params  json.RawMessage
// }

// // Fetches all features from the database
// func getFeatures() ([]Feature, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, name FROM features")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var features []Feature
// 	for rows.Next() {
// 		var f Feature
// 		if err := rows.Scan(&f.ID, &f.Name); err != nil {
// 			return nil, err
// 		}
// 		features = append(features, f)
// 	}
// 	return features, nil
// }

// // Fetches all operations for a given feature ID
// func getOperationsByFeatureID(fid int) ([]Operation, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, name, f_id FROM operations WHERE f_id = ?", fid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var ops []Operation
// 	for rows.Next() {
// 		var op Operation
// 		if err := rows.Scan(&op.ID, &op.Name, &op.FID); err != nil {
// 			return nil, err
// 		}
// 		ops = append(ops, op)
// 	}
// 	return ops, nil
// }

// // Fetches all variations for a given operation ID
// func getVariationsByOperationID(oid int) ([]Variation, error) {
// 	db, err := connectDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, name, o_id, formula, params FROM variations WHERE o_id = ?", oid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var variations []Variation
// 	for rows.Next() {
// 		var v Variation
// 		if err := rows.Scan(&v.ID, &v.Name, &v.OID, &v.Formula, &v.Params); err != nil {
// 			return nil, err
// // javascript
// Copy
// Edit
		}
// 		variations = append(variations, v)
// 	}
// 	return variations, nil
// }

// Step-by-Step Flow
// 🔷 1. Display Feature Menu
// You run the program.

// It connects to the da// markdown
// Copy
// Edit
tabase and fetches all entries from the features table.

// Example:

// javascript
// Copy
// Edit
// 1. Math
// 2. String
// 🔷 2. User Selects a Feature
// Suppose you choose 2 (String).

// The program fetches all operations with f_id = 2 from the operations table.

// Example:

// markdown
// Copy
// Edit
// 5. Palindrome
// 6. Reverse String
// 7. Count Vowels
// 🔷 3. User Selects an Operation
// You select an operation (e.g., 7 Count Vowels).

// The program fetches all variations under that operation from the variations table usin// css
// Copy
// Edit
g o_id.

// 🔷 4. Select a Variation
// If there's only one variation, it's auto-selected.

// If multiple, you're asked to choose one.

// The variation includes:

// formula: e.g., vowelcount(input)

// params: e.g., {"input": "string"}

// 🔷 5. Provide Input
// The program reads the params, prompts th// go
// Copy
// Edit
e user:

// css
// Copy
// Edit
// Enter value for input: Kiran
// Collects inputs in a map and sends it to evaluateFormula() function.

// 🔷 6. Evaluate the Formula
// The formula string like vowelcount(input) is pa// makefile
// Copy
// Edit
rsed using the govaluate library.

// Custom functions (like vowelcount, reverse) are registered with govaluate.

// Example:
// go
// Copy
// Edit
// "vowelcount": func(args ...interface{}) (interface{}, error) {
//   // Count vowels in a string
// }
// The formula is dynamically evaluated using the input values.

// 🔷 7. Print Result
// The result is printed:

// makefile
// Copy
// Edit
// Result: 2
// 🔷 8. Log the Operation
// Finally, the program inserts a record into the user_logs table:

// category (e.g., String)

// program (e.g., Count Vowels)

// variation_id

// inputs (as JSON)

// result

// status*/
