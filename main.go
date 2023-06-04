package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var num float64 = math.Floor(3.14)
	var someString string = "some string"
	var titleString string = strings.ToUpper(someString)

	fmt.Printf("The floor of 3.14 is %v\n", num)
	fmt.Println(titleString)
}