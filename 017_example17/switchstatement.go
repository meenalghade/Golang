package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("the ouput is printed")
		//fallthrough
	case false:
		fmt.Println("Case not executed")
		//fallthrough
	case (2 == 2):
		fmt.Println("This is a true statement")
		fallthrough
	case (2 != 2):
		fmt.Println("this is false statement")
		fallthrough
	case false:
		fmt.Println("This a false statement")
	case (14 != 14):
		fmt.Println("this is treu statement")
	}
}
