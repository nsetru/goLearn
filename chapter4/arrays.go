package main

import "fmt"

func main() {
	x := [5]float64{10, 20, 30, 40, 50}
	// x[0] = 10
	// x[1] = 20
	// x[2] = 30
	// x[3] = 40
	// x[4] = 50

	// Different ways of looping through array
	// Method 1: Use fixed length 5
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println("Method 1: Total", total/5)

	// Method 2: Use `len` to get length of an array
	var total1 float64 = 0
	for i := 0; i < len(x); i++ {
		total1 += x[i]
	}
	fmt.Println("Method 2: Total", total1/float64(len(x)))

	// Method 3: Use `range` keyword for looping
	var total2 float64 = 0
	for _, value := range x {
		total2 += value
	}
	fmt.Println("Method 3: Toatl", total2/float64(len(x)))
}
