package main

import "fmt"

type person struct {
	first  string
	last   string
	age    int
	salary float32
}

func main() {
	p1 := person{
		first:  "Meenal",
		last:   "Ghade",
		age:    37,
		salary: 70000.00,
	}
	p2 := person{
		first:  "NIkhil",
		last:   "Ghade",
		age:    36,
		salary: 120000.00,
	}
	p3 := person{
		first:  "Vivaan",
		last:   "Ghade",
		age:    5,
		salary: 10000.00,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
