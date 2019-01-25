package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a)
	b, c := bar()
	fmt.Println(b, c)
}

func foo() int {
	return 45
}

func bar() (int, string) {
	return 1942, "i M WATCHING BIG BOSS"
}
