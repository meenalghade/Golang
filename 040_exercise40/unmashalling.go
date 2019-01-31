package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"Mike","Last":"Ross","Age":28},{"First":"Harvey","Last":"Specter","Age":35}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data ", people)

	for i, val := range people {
		fmt.Println("Person Number", i+1)
		fmt.Println(val.First, val.Last, val.Age)
	}
}
