package main

import "fmt"

func shapeSubMenu(shape string) {
	for {
		fmt.Printf("\n%s:\n", shape)
		fmt.Println("1. Area")
		fmt.Println("2. Volume")
		fmt.Println("3. Back")

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
			switch shape {
			case "Triangle":
				areaOfTriangle()
			case "Square":
				areaOfSquare()
			case "Rectangle":
				areaOfRectangle()
			case "Circle":
				areaOfCircle()
			}
		case 2:
			switch shape {
			case "Triangle":
				volumeOfTriangle()
			case "Square":
				volumeOfSquare()
			case "Rectangle":
				volumeOfRectangle()
			case "Circle":
				volumeOfCircle()
			}
		case 3:
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
