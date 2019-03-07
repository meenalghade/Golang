package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
	Age     int
}

func main() {
	p1 := person{
		First:   "Harvey",
		Last:    "Specter",
		Sayings: []string{"Shaken by the words", "Any Last Wishes", "Dont last for ever"},
		Age:     0,
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, er := json.Marshal(a)
	if er != nil {
		return []byte{}, fmt.Errorf("There is an error in toJSON: %v", er)
	}
	return bs, nil
}
