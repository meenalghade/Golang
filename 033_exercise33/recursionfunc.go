package main

import "fmt"

func main() {
	val := factorial(8)
	fmt.Println(val)

	fac := loopfactorial(8)
	fmt.Println(fac)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopfactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
