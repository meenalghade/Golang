package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mysum(2, 3))
	fmt.Println("9 + 8 =", mysum(9, 8))
	fmt.Println("10 + 20 =", mysum(10, 20))
}

func mysum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
