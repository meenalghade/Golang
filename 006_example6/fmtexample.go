package main 

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n",y)    // type of y vale 
	fmt.Printf("%b\n",y) 	// the value of y in in binary format
	fmt.Printf("%x\n",y)	//value of y in decimal format
	fmt.Print("%#x\n",y)	//value of y in hexadecimal format

	y=698
	fmt.Printf("%#x\n",y)
	fmt.Printf("%#x\t%b\t%x\n",y,y,y)

	s:=fmt.Sprintf("%#x\t%b\t%x\n",y,y,y)
	fmt.Println(s)
	fmt.Printf("%v",y) 
}
/*Output of the program is 
42
int
101010
2a
%#x
420x2ba
0x2ba   1010111010      2ba
0x2ba   1010111010      2ba

698
*/