package main

import (
	"fmt"
	"sort"
)

//Person type added
type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := Person{"James", 34}
	p2 := Person{"April", 29}
	p3 := Person{"Diana", 12}
	p4 := Person{"Zane", 44}
	p5 := Person{"Freed", 56}

	people := []Person{p1, p2, p3, p4, p5}

	fmt.Println("Before Sorting by age", people)
	sort.Sort(ByAge(people))
	fmt.Println("After Sorting by Age ", people)

	fmt.Println("Before Sorting by Name", people)
	sort.Sort(ByName(people))
	fmt.Println("After Sorting by Name ", people)

}
