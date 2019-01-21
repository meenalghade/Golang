package main

import "fmt"

func main() {
	foo()
	defer foo()
	bar()

}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("I m in bar")
}
