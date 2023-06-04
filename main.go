package main

import (
	"math"
	"strings"
	"fmt"
)

func main() {
	var num float64 = math.Floor(3.14)
	var titleString string = strings.ToTitle("some string")

	fmt.Printf("The floor of 3.14 is %v\n", num)
	fmt.Println(titleString)
}