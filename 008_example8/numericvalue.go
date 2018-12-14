package main

import "fmt"
import "runtime"

func main(){
	z := 76
	y := 34.877
	fmt.Println(z)
	fmt.Println(y)
	fmt.Printf("%T\n",z)
	fmt.Printf("%T\n",y)

	fmt.Printf(runtime.GOLANGOS)
	fmt.Printf(runtime.GOARCHITE)   
}

//76
//34.877
//int
//float64
