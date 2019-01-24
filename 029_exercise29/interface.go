package main

import (
	"fmt"
)

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

func (p person) speak() {
	fmt.Println("I am ", p.firstname, p.lastname)
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am passed into bar ", h.(person).firstname)
	case secreteAgent:
		fmt.Println("I am passed into bar ", h.(secreteAgent).firstname)
	}

	fmt.Println("I am passed into bar ", h)
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

	p1 := person{
		firstname: "Dr.",
		lastname:  "YES",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)
	p1.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
}
