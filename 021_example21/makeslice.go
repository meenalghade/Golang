package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 33
	x[6] = 44
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 22, 44, 88, 99, 46)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
