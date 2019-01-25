package main

import "fmt"

func main() {

	a := limitvalue()
	b := limitvalue()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func limitvalue() func() int {
	var x int
	return func() int {
		x++
		return x
	}

}
