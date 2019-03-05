package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receiver
	cs := make(chan<- int) //send

	fmt.Println("----------------------------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	//specific to general doent assign
	//c = cr    /cannot use cr(type <-chan int) as type chan int in assignment
	//c = cs    /cannot use cs(type chan<- int) as type chan int in assignment

	//general to specific converts
	fmt.Println("----------------------------------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	//speciifcs to general doesn't converts
	fmt.Println("----------------------------------")
	//fmt.Printf("c\t%T\n", (chan int)(cr)) //cannot convert cr (type <-chan int) to type chan int go
	//fmt.Printf("c\t%T\n", (chan int)(cs)) //cannot convert cs(type chan<- int) to type chan int go
}
