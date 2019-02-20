package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(" GO OS", runtime.GOOS)
	fmt.Println("GO ARCH ", runtime.GOARCH)
	fmt.Println("CPU'S", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())

	go foo()
	bar()

	fmt.Println("CPU'S", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo :", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar  :", i)
	}
}
