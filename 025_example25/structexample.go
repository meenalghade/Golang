package main

import "fmt"

type person struct {
	first  string
	last   string
	age    int
	salary float32
}

type deparment struct {
	person
	role       string
	department string
}

func main() {
	//struct example
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

	//embedded struct
	p4 := deparment{
		person: person{
			first:  "Sudhir",
			last:   "Ghade",
			age:    65,
			salary: 80000.00,
		},
		role:       "Head of the family",
		department: "Mechanical Engineering",
	}
	//Anonymous structs
	p5 := struct {
		first string
		last  string
		age   int
	}{
		first: "Smita",
		last:  "Ghade",
		age:   59,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)
}
