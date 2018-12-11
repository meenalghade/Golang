package main 

import "fmt"

type cheese int 
var x cheese
var y int 
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42 
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)
}
//0
//main.cheese
//42
//42
//int
