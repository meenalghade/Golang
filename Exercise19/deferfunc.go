package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello World")
}

func foo() {
	defer func() {
		fmt.Println("I am in the defer foo()")
	}()
	fmt.Println(" I m in function foo()")
}
