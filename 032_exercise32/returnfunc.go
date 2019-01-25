package main

import "fmt"

func main() {
	r := foo()
	fmt.Println(r)

	s := bar()
	fmt.Printf("%T\n", s)

	t := s()
	fmt.Println(t)

	fmt.Println("The return func value of a func bar is :", bar()())
}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
