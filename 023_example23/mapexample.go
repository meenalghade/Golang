package main

import "fmt"

func main() {
	m := map[string]int{"Sandeep": 32, "Aarti": 23, "Ashwini": 31}
	fmt.Println(m)

	fmt.Println(m["Sandeep"])
	fmt.Println(m["Amol"])
	//checking the value in map
	v, ok := m["Amol"]
	fmt.Println(v)
	fmt.Println(ok)
	// or other way of iterating the map is
	if v, ok := m["Ashwini"]; ok {
		fmt.Println("This is value of Ashwini:", v)
	}

	m["Amol"] = 34

	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "Aarti")
	fmt.Println(m)

}
