package main

import (
	"fmt"
	"math"
	"strings"
)

func square(num float64) float64 {
	// Formula -> num * num
	return num * num
}

func area(radius float64) float64 {
	// Formula -> PI * radius * radius
	return math.Round(math.Pi * square(radius) * 100) / 100
}

func main() {
	// Testing
	var num float64 = math.Floor(3.14)
	var someString string = "some string"
	var titleString string = strings.ToUpper(someString)

	fmt.Print("Enter radius: ")
	var inputNumber int
	_, err := fmt.Scan(&inputNumber)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var computedArea = area(float64(inputNumber))

	fmt.Printf("The floor of 3.14 is %v\n", num)
	fmt.Printf("The square of %v is %v\n", num, square(num))
	fmt.Printf("The area of a circle of radius %v units is %v units\n", inputNumber, computedArea)

	fmt.Println()
	fmt.Println(titleString)
}