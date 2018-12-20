package main

import "fmt"

func main() {
	m := map[string][]string{
		"Janes Bond":      []string{`Shaken,not stirred`, `Martin`, `Women`},
		"Moneypenny_miss": []string{`James Bond`, `literature`, `Computer Science`},
		"no_dr":           []string{`Internet`, `being evil`, `sunset`},
	}

	m["SHahrukh Khan"] = []string{"deewana", "dilashiyana", "Bazigar"}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("The value of record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
