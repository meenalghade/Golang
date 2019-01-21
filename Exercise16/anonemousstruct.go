package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Vivaan",
		friends: map[string]int{
			"Pihu":   333,
			"armaan": 777,
			"Aliya":  999,
		},
		favDrinks: []string{
			"pepsi",
			"sprite",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
