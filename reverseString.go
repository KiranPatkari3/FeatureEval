package main

import (
	"fmt"
)

func reverseString() {
	var str string
	fmt.Print("Enter any string: ")
	fmt.Scan(&str)
	original := str
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	result := reversed
	status := "Success"
	fmt.Println("Reversed String:", result)

	writeLog("String", "Reverse String", original, result, status)
	writeDBLog("String", "Reverse String", fmt.Sprintf("input=%s", str), result, "Success")

}
