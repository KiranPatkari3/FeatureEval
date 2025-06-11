package main

import (
	"fmt"
	"strings"
	"unicode"
)

func checkPalindrome() {
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	// Check if input contains only letters
	for _, ch := range str {
		if !unicode.IsLetter(ch) {
			fmt.Println("Invalid input! Please enter only alphabets.")
			return
		}
	}

	// Convert to lowercase to make it case-insensitive
	str = strings.ToLower(str)

	// Reverse the string
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	var result, status string
	if str == reversed {
		result = "It's a palindrome"
		status = "Success"
	} else {
		result = "Not a palindrome"
		status = "Failed"
	}

	fmt.Println(result)
	writeLog("String", "Palindrome Check", str, result, status)
	writeDBLog("String", "Check Palindrome", fmt.Sprintf("input=%s", str), result, "Success")

}
