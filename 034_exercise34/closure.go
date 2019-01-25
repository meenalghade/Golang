package main

import "fmt"

//var x int
func main() {

	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	x--
	///code block within the braces
	{
		y := 8
		fmt.Println(y)
	}
	//fmt.Println(y)
	fmt.Println(x)
	foo()
	fmt.Println(x)
}
func foo() {
	fmt.Println("Hello World")
	//x++
}
