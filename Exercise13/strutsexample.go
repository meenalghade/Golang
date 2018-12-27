package main

import "fmt"

type person struct {
	first       string
	last        string
	favIcecream []string
}

func main() {
	p1 := person{
		first: "Shruti",
		last:  "Mundhada",
		favIcecream: []string{
			"Stawberry",
			"vanilla",
			"choclate",
		},
	}

	p2 := person{
		first: "Manajri",
		last:  "Rajput",
		favIcecream: []string{
			"kulfi",
			"falooda",
			"almond",
		},
	}

	p3 := person{
		first: "Meenal",
		last:  "Ghade",
		favIcecream: []string{
			"Sitafal",
			"Chocobar",
			"Rasberry",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
		p3.last: p3,
	}
	fmt.Println(m)

	//other way to print the value

	for k, v := range m {
		fmt.Print(k)
		fmt.Println(v.first, "     ", v.last)
		for i, val := range v.favIcecream {
			fmt.Println(i, val)
		}
	}

	//fmt.Println(p1.first, "     ", p1.last)
	//for i, v := range p1.favIcecream {
	//	fmt.Println(i, v)
	//}

	//fmt.Println(p2.first, "    ", p2.last)
	//for i, v := range p2.favIcecream {
	//	fmt.Println(i, v)
	//}

	//fmt.Println(p3.first, "    ", p3.last)
	//for i, v := range p3.favIcecream {
	//	fmt.Println(i, v)
	//}
}
