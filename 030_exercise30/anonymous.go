package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I m in the anonymous function")
	}()

	func(x int) {
		fmt.Println("The value of x is :", x)
	}(42)
}
