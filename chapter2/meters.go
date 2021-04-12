// Write program that converts from feet into meters. (1 ft = 0.3048 m)
package main

import "fmt"

func main() {
	fmt.Print("Enter the measurement in feet : ")
	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.3048
	fmt.Println("Measurement in meters = ", meters)
}
