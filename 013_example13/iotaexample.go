package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	/*kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb*/

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}

//2               10
//4               100
//1024            10000000000
//1048576         100000000000000000000    increase by 10 zeros
//1073741824      1000000000000000000000000000000    increased by 10 zeros
