package main 

import "fmt"

var x string
var y int 
var z float64
var b bool
var a *int
var c [10]int
var d [10]string
func main() {
    
	fmt.Println(x)
	fmt.Printf("%T\n",x)	

	fmt.Println(y)
	fmt.Printf("%T\n",y)

	fmt.Println(z)
	fmt.Printf("%T\n",z)

	fmt.Println(b)
	fmt.Printf("%T\n",b)

	/*Output of above code is 

	string
	0
	int
	0
	float64
	false
	bool
	*/
	fmt.Println(a)
	fmt.Printf("%T\n",a)

	fmt.Println(c)
	fmt.Printf("%T\n",c)

	fmt.Println(d)
	fmt.Printf("%T\n",d)

	/* output of above code 
	<nil>
	*int
	[0 0 0 0 0 0 0 0 0 0]
	[10]int
	[         ]
	[10]string
	*/
}