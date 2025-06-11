package main

import (
	"fmt"
)

func areaOfTriangle() {
	var base, height float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter base of triangle: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter base of triangle: ")
		}

		_, err = fmt.Scan(&base)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	firstTry = false
	for {
		if !firstTry {
			fmt.Print("Enter height of triangle: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter height of triangle: ")
		}

		_, err = fmt.Scan(&height)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear)
			continue
		}
		break
	}

	area := 0.5 * base * height
	fmt.Println("Area of Triangle:", area)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", base, height)
	result := fmt.Sprintf("%.2f", area)
	writeLog("Math", "Area of Triangle", inputs, result, "Success")
	writeDBLog("Math", "Area of Triangle", fmt.Sprintf("base=%.2f, height=%.2f", base, height), fmt.Sprintf("%.2f", area), "Success")

}
func volumeOfTriangle() {
	var base, height, length float64
	var err error

	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter base of triangle: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter base of triangle: ")
		}

		_, err = fmt.Scan(&base)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	firstTry = false

	for {
		if !firstTry {
			fmt.Print("Enter height of triangle:  ")
			firstTry = true
		} else {
			fmt.Print("ReEnter height of triangle: ")
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
	firstTry = false
	for {
		if !firstTry {
			fmt.Print("Enter length of triangle:  ")
			firstTry = true
		} else {
			fmt.Print("ReEnter length of triangle: ")
		}

		_, err = fmt.Scan(&length)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear) // clear invalid input
			continue
		}
		break
	}

	volume := 0.5 * base * height * length
	fmt.Println("Volume of triangular :", volume)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", base, height)
	result := fmt.Sprintf("%.2f", volume)
	writeLog("Math", "Volume of Triangle", inputs, result, "Success")
	writeDBLog("Math", "Volume of Triangle", fmt.Sprintf("base=%.2f, height=%.2f", base, height, length), fmt.Sprintf("%.2f", volume), "Success")

}
