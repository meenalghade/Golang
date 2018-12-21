package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavour []string
}

func main() {
	p1 := person{
		first: "hello",
		last:  "world",
		favFlavour: []string{
			"choclate",
			"vanilla",
		},
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavour {
		fmt.Println(i, v)
	}
}
