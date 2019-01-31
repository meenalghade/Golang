package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	// '&' ampercent will give the address of the location that is called a pointer
	fmt.Println(&a)
	// to the type of the a
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	//assinging address to a variable
	// '*' with type is pointer to that type
	b := &a
	fmt.Println(b)
	fmt.Println(*b)  // to get the value at address b we use '*b'
	fmt.Println(*&a) // this will be taking the value from the address location

	*b = 44
	fmt.Println(a) // here the value of a gets changed to the value of b a both are pointing to the same location
}
