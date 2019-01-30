package main

import "fmt"

func main() {
	val := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(val...)
	fmt.Println(sum)

	sum1 := bar(val)
	fmt.Println(sum1)

}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}
