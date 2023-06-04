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

	fmt.Printf("The floor of 3.14 is %v\n", num)
	fmt.Printf("The square of %v is %v\n", num, square(num))
	fmt.Printf("The area of a circle of radius 5 units is %v units\n", area(5))

	fmt.Println()
	fmt.Println(titleString)
}