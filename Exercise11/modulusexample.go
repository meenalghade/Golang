package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 != 0 {
			fmt.Printf("the remainder of %d divide by 4 is %d\n", i, i%4)
		} else {
			fmt.Printf("the remainder of %d divide by 4 is %d\n", i, i%4)
		}
	}
}
