package main

import "fmt"

func areaOfRectangle() {
	var length, width float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter length of triangle: ")
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

	firstTry = false
	for {
		if !firstTry {
			fmt.Print("Enter width of triangle: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter width of triangle: ")
		}

		_, err = fmt.Scan(&width)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear)
			continue
		}
		break
	}

	area := length * width
	fmt.Println("Area of Rectangle:", area)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", length, width)
	result := fmt.Sprintf("%.2f", area)
	writeLog("Math", "Area of Triangle", inputs, result, "Success")
	writeDBLog("Math", "Area of Triangle", fmt.Sprintf("base=%.2f, height=%.2f", length, width), fmt.Sprintf("%.2f", area), "Success")

}
func volumeOfRectangle() {
	var length, width, height float64
	var err error
	firstTry := false

	for {
		if !firstTry {
			fmt.Print("Enter length of triangle: ")
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

	firstTry = false
	for {
		if !firstTry {
			fmt.Print("Enter width of triangle: ")
			firstTry = true
		} else {
			fmt.Print("ReEnter width of triangle: ")
		}

		_, err = fmt.Scan(&width)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			var clear string
			fmt.Scanln(&clear)
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

	volume := length * width * height
	fmt.Println("Volume of cuboid:", volume)
	inputs := fmt.Sprintf("base=%.2f, height=%.2f", length, width, height)
	result := fmt.Sprintf("%.2f", volume)
	writeLog("Math", "Volume of Rectangle", inputs, result, "Success")
	writeDBLog("Math", "Volume of Rectangle", fmt.Sprintf("base=%.2f, height=%.2f", length, width, height), fmt.Sprintf("%.2f", volume), "Success")

}
