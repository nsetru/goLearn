// Using the example program as a starting point, write a program that
// converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)

package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in Fahrenheit")
	var input int
	fmt.Scanf("%d", &input)

	celsius := (input - 32) * 5 / 9
	fmt.Println("Celsius = ", celsius)
}
