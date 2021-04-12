package main

import "fmt"

// Slices
// var x []float64
func main() {
	arr := [5]float64{1, 2, 3, 4, 5}
	// arr[low:high]
	x := arr[0:5]
	fmt.Println("1. array slice [0:5]", x)

	//For convenience we are also allowed to omit low,
	//high or even both low and high

	// Omit low
	y1 := arr[:3]
	fmt.Println("2. array slice [:3]", y1)

	// Omit high
	y2 := arr[1:]
	fmt.Println("3. array slice[1:]", y2)

	// Omit both high and low
	y3 := arr[:]
	fmt.Println("4. array slicep[:]", y3)

	// High is length of array
	y4 := arr[0:len(arr)]
	fmt.Println("5. array[0:len(arr)]", y4)

	// append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Before append:", slice1)
	fmt.Println("After append:", slice2)

	// copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("Before copy", slice3)
	fmt.Println("After copy", slice4)

}
