package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countVowels() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	// Validate: input must contain only letters
	for _, ch := range input {
		if !unicode.IsLetter(ch) {
			fmt.Println("Invalid input! Please enter only alphabets.")
			return
		}
	}

	// Convert to lowercase to handle both uppercase and lowercase
	input = strings.ToLower(input)

	vowelCount := 0
	for _, ch := range input {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			vowelCount++
		}
	}

	result := fmt.Sprintf("Vowel Count: %d", vowelCount)
	status := "Success"
	fmt.Println(result)

	writeLog("String", "Count Vowels", input, result, status)
	writeDBLog("String", "Check Palindrome", fmt.Sprintf("input=%s", input), result, "Success")

}
