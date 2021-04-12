package main

import "fmt"

func main() {
	count := 100
	for i := 1; i <= count; i++ {
		if i%3 == 0 {
			fmt.Println("Divisible by 3 = ", i)
		}
	}
}
