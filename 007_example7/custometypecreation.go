package main 

import "fmt"

var a int 

type hotdog int 

var b hotdog

func main() {
	a= 43
	b= 42
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	fmt.Println(b)
	fmt.Printf("%T\n",b)
	// a=b	.\custometypecreation.go:18:4: cannot use b (type hotdog) as type int in assignment
	a = int(b)    // this the conversion of one type to other type in go programming 
	fmt.Println(a)   // there is nothing called as type casting in go programming 
	fmt.Printf("%T\n",a) 
}
/* Output of above program 

43
int
42
main.hotdog
42
int
*/