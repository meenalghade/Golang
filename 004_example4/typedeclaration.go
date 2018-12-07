package main 

import "fmt"

var y = 42
//declare a variable with the identifier "z" is of Type String 	
//and assign the value "Shaken, not sirred "
var z string ="Shaken, not stirred"
//this is the STATIC programming language 
// a VARIABLE is DECLARED to hold the VALUE of certain TYPE
// not a DYNAMIC programming Language
var a string =`Meenal said, "Shaken, not stirred"`
//other way to use string is by useing the backqoute '`' to use the '"' in a string value 
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n",y)
	fmt.Println(z)
	fmt.Printf("%T\n",z)
	//z=46  Error :.\typedeclaration.go:16:4: cannot use 46 (type int) as type string in assignment
	z = "Hello type declaration"
	fmt.Println(z)
	fmt.Printf("%T\n",z)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
}