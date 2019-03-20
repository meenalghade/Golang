package word

import (
	"fmt"
	"testing"

	"github.com/meenalghade/Golang/063_Example_63/qoute"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "Expected ", 3)
	}
}

func TestUserCount(t *testing.T) {
	n := UserCount("one two three three three")
	for k, v := range n {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got ", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

//this is coverage function to check count
func ExampleCount() {
	fmt.Println(Count("one two three"))
	//Output:
	//3
}
func BenchmarkUserCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserCount(qoute.SumAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(qoute.SumAlso)
	}
}
