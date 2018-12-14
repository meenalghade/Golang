package main

import "fmt"

func main() {
	s := "Hello GO Programming"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i:=0; i < len(s); i++{
		fmt.Printf("%#U ",s[i])  //UTF8 representation  format
	}
	fmt.Println("")
	for i,v := range s {
		fmt.Printf("At positon %d we have hex %#U\n",i,v)
	}
}
//Hello GO Programming
//string
//[72 101 108 108 111 32 71 79 32 80 114 111 103 114 97 109 109 105 110 103]
//[]uint8

//U+0048 'H' U+0065 'e' U+006C 'l' U+006C 'l' U+006F 'o' U+0020 ' ' U+0047 'G' U+004F 'O' U+0020 ' ' U+0050 'P' U+0072 'r' U+006F 'o' U+0067 'g' U+0072 'r' U+0061 'a' U+006D 'm' U+006D 'm' U+0069 'i' U+006E 'n' U+0067 'g'

// At positon 0 we have hex U+0048 'H'
//At positon 1 we have hex U+0065 'e'
//At positon 2 we have hex U+006C 'l'
//At positon 3 we have hex U+006C 'l'
//At positon 4 we have hex U+006F 'o'
//At positon 5 we have hex U+0020 ' '
//At positon 6 we have hex U+0047 'G'
//At positon 7 we have hex U+004F 'O'
///At positon 8 we have hex U+0020 ' '
//At positon 9 we have hex U+0050 'P'
//At positon 10 we have hex U+0072 'r'
//At positon 11 we have hex U+006F 'o'
//At positon 12 we have hex U+0067 'g'
//At positon 13 we have hex U+0072 'r'
//At positon 14 we have hex U+0061 'a'
//At positon 15 we have hex U+006D 'm'
//At positon 16 we have hex U+006D 'm'
//At positon 17 we have hex U+0069 'i'
//At positon 18 we have hex U+006E 'n'
//At positon 19 we have hex U+0067 'g'


