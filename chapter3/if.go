package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("even = ", i)
		} else {
			fmt.Println("odd = ", i)
		}
	}
}
