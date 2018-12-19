package main

import "fmt"

const (
	a = 2017 + iota
	b
	c
)

const (
	x = 2017 + iota
	y = 2017 + iota
	z = 2017 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

//2017
//2018
//2019
//2017
//2018
//2019
