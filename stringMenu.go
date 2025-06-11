package main

import "fmt"

func stringMenu() {
	for {
		fmt.Println("\nString Programs:")
		fmt.Println("1. Check Palindrome")
		fmt.Println("2. Reverse String")
		fmt.Println("3. Count Vowels")
		fmt.Println("4. Back to Main Menu")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, try again.")
			var clear string
			fmt.Scanln(&clear)
			continue
		}

		switch choice {
		case 1:
			checkPalindrome()
		case 2:
			reverseString()
		case 3:
			countVowels()
		case 4:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
