package main

import (
	"fmt"
	"strings"

	mystring "github.com/meenalghade/Golang/benchmarkExample/stringfunc"
)

const s = "Two touch the string, The harp is dumb. The strings to our past are burned because we need a new beginning. Strings of gravity vibrate at a different frequency than strings of light. We are nothing but a string of gut on a stick of bone riding this piece of astral soot for one piteous splinter of eternity."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Println("\n\n", mystring.Cat(xs))
	fmt.Println("\n\n", mystring.Join(xs))
}
