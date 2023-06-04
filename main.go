package main

import (
	"fmt"
	"math"
	"strings"
)

func square(num float64) float64 {
	return num * num
}

func main() {
	var num float64 = math.Floor(3.14)
	var someString string = "some string"
	var titleString string = strings.ToUpper(someString)

	fmt.Printf("The floor of 3.14 is %v\n", num)
	fmt.Printf("The square of %v is %v\n", num, square(num))
	fmt.Println(titleString)
}