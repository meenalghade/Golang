package main

import "fmt"

func main() {
	jb := []string{"Parul", "Paul", "Chocalate", "Vanilla", "Buttersoch"}
	fmt.Println(jb)
	dh := []string{"Diana", "Heden", "Vanilla", "Chocochip", "Rasberry"}
	fmt.Println(dh)

	x := [][]string{jb, dh}
	fmt.Println(x)

	for i, s := range x {
		fmt.Println("record  :", i)
		for j, val := range s {
			fmt.Printf("\t indexposition : %v\t  value: %v\n", j, val)
		}
	}
}
