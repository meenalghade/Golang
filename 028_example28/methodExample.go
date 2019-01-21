package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secreteAgent struct {
	person
	ltk bool
}

func (s secreteAgent) speak() {
	fmt.Println("I am ", s.firstname, s.lastname)
}

func main() {
	sa1 := secreteAgent{
		person: person{
			"Julie",
			"Robert",
		},
		ltk: true,
	}

	sa2 := secreteAgent{
		person: person{
			"Angolina",
			"Jolie",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

}
