package main

import "fmt"

func mathMenu() {

	for {
		fmt.Println("\nChoose a shape:")
		fmt.Println("1. Triangle")
		fmt.Println("2. Square")
		fmt.Println("3. Rectangle")
		fmt.Println("4. Circle")
		fmt.Println("5. Back")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input.")
			var clear string
			fmt.Scanln(&clear)
			continue
		}

		switch choice {
		case 1:
			shapeSubMenu("Triangle")
		case 2:
			shapeSubMenu("Square")
		case 3:
			shapeSubMenu("Rectangle")
		case 4:
			shapeSubMenu("Circle")
		case 5:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
