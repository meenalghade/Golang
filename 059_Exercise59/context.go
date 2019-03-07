package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("the context is\t", ctx)
	fmt.Println("the context err\t", ctx.Err())
	fmt.Printf("the context type is\t%T\n", ctx)
	fmt.Println("-------------------------------------------")
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("the context is\t", ctx)
	fmt.Println("the context err\t", ctx.Err())
	fmt.Printf("the context type is\t%T\n", ctx)

	fmt.Println("the cancel\t", cancel)
	fmt.Printf("the cancel type is\t%T\n", cancel)
	fmt.Println("-------------------------------------------")
	cancel()
	fmt.Println("the context is\t", ctx)
	fmt.Println("the context err\t", ctx.Err())
	fmt.Printf("the context type is\t%T\n", ctx)

	fmt.Println("the  cancel\t", cancel)
	fmt.Printf("the cancel type is\t%T\n", cancel)

	fmt.Println("-------------------------------------------")

}
