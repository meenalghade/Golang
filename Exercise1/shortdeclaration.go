package main

import "fmt"


func main() {
	x := 42
	fmt.Println(x)
	y := "James Bond"
	fmt.Println(y)
	z := true
	fmt.Println(z)
	fmt.Printf("%v\t %v \t %v\t",x,y,z)
}
//$ go run shortdeclaration.go
//42
//James Bond
//true
//42       James Bond      true
