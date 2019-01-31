package main

import "fmt"

func main() {
	x := 0
	fmt.Println("value of x before", x)
	fmt.Println("value of x before", &x)
	foo(&x)
	fmt.Println("value of x after", x)
	fmt.Println("value of x after", &x)
}

func foo(y *int) {
	fmt.Println("value of y before", y)
	fmt.Println("value of y before", *y)

	*y = 55
	fmt.Println("value of y after", y)
	fmt.Println("value of y after", *y)

}
