package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}

	// assigning a value to the function
	p := func(x int) {
		fmt.Println("The love story during independence is ", x)
	}

	f()
	p(1942)
}
