package main

import "fmt"

func main() {
	fmt.Print("Enter the number")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
