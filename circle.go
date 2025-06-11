package main

import (
	"fmt"
	"math"
)

func areaOfCircle() {
	var radius float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter radius of circle: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter radius of circle: ")
		}

		_, err = fmt.Scan(&radius)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	area := math.Pi * radius * radius
	fmt.Println("Area of Circle:", area)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", radius)
	result := fmt.Sprintf("%.2f", area)
	writeLog("Math", "Area of Circle", inputs, result, "Success")
	writeDBLog("Math", "Area of Circle", fmt.Sprintf("base=%.2f, height=%.2f", radius), fmt.Sprintf("%.2f", area), "Success")
}
func volumeOfCircle() {
	var radius, height float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter radius of cylinder: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter radius of cylinder:")
		}

		_, err = fmt.Scan(&radius)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	for {
		if !firstTry {
			fmt.Print("Enter height of cylinder: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter height of cylinder:")
		}

		_, err = fmt.Scan(&height)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	volume := math.Pi * radius * radius * height
	fmt.Println("Volume of cylinder:", volume)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", radius, height)
	result := fmt.Sprintf("%.2f", volume)
	writeLog("Math", "Volume of Circle", inputs, result, "Success")
	writeDBLog("Math", "Volume of Circle", fmt.Sprintf("base=%.2f, height=%.2f", radius, height), fmt.Sprintf("%.2f", volume), "Success")

}
