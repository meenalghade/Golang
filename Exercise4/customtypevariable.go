package main 

import "fmt"

type cheese int 
var x cheese
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42 
	fmt.Println(x)
}
//0
//main.cheese
//42
