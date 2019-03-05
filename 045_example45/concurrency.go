package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(" GO OS", runtime.GOOS)
	fmt.Println("GO ARCH ", runtime.GOARCH)
	fmt.Println("CPU'S", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	wg.Wait()

	fmt.Println("CPU'S", runtime.NumCPU())
	fmt.Println("GO Routines", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo :", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar  :", i)
	}
}
