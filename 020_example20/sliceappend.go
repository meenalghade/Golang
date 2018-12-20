package main

import "fmt"

func main() {
	x := []int{21, 32, 76, 88, 23, 65, 33}
	fmt.Println(x)
	x = append(x, 43, 90, 12)
	fmt.Println(x)

	y := []int{31, 76, 33, 99, 93, 54}
	fmt.Println(y)
	y = append(x, y...)
	fmt.Println(y)

	x = append(x[:2], x[4:]...) //delete from slice
	fmt.Println(x)

}
