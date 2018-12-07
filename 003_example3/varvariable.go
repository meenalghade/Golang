package main

import "fmt"

var y = 34

//Declare there is a variable with the identifier "z"
// and that the variable with the Identifier "z" is of type int
// Assigns the zero valueof type int to "z"
// false for booleans, 0 for intergers, 0.0 for floats,
// "" for Strings and nil for pointers,functions,interfaces,slices,channels and maps.
var z int

func main() {
	x := 55
	fmt.Println(x)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("the value of y :", y)
}
