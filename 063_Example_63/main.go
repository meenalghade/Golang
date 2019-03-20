package main

import (
	"fmt"

	"github.com/meenalghade/Golang/063_Example_63/qoute"
	"github.com/meenalghade/Golang/063_Example_63/word"
)

func main() {
	fmt.Println(word.Count(qoute.SumAlso))

	for k, v := range word.UserCount(qoute.SumAlso) {
		fmt.Println(v, k)
	}
}

/* Running steps :
go run main.go
cd word/
go test
go test -bench .
go test -coverprofile c.out
go tool cover -html=c.out
godoc -http=:8282




*/
