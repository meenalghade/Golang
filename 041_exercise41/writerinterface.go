package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Go Lang")
	fmt.Fprintln(os.Stdout, "Hello Go Lang")
	io.WriteString(os.Stdout, "Hello World")
}
