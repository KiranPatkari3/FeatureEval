package main

import "fmt"

func areaOfSquare() {
	var side float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter side of square: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter side of square: ")
		}

		_, err = fmt.Scan(&side)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}
	area := side * side
	fmt.Println("Area of Square:", area)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", side)
	result := fmt.Sprintf("%.2f", area)
	writeLog("Math", "Area of Triangle", inputs, result, "Success")
	writeDBLog("Math", "Area of Triangle", fmt.Sprintf("side=%.2f", side), fmt.Sprintf("%.2f", area), "Success")

}
func volumeOfSquare() {
	var side float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter side of square: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter side of square: ")
		}

		_, err = fmt.Scan(&side)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	volume := side * side * side
	fmt.Println("Volume of cube:", volume)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", side)
	result := fmt.Sprintf("%.2f", volume)
	writeLog("Math", "Volume of Square", inputs, result, "Success")
	writeDBLog("Math", "Volume of Square", fmt.Sprintf("base=%.2f, height=%.2f", side), fmt.Sprintf("%.2f", volume), "Success")

}
