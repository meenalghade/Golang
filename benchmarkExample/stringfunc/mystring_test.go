package mystring

import (
	"fmt"
	"strings"
	"testing"
)

const c = "Two touch the string, The harp is dumb. The strings to our past are burned because we need a new beginning. Strings of gravity vibrate at a different frequency than strings of light. We are nothing but a string of gut on a stick of bone riding this piece of astral soot for one piteous splinter of eternity."

var xs []string

/*func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "wants", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "wants", "Shaken not stirred")
	}
}*/

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	//Output:
	//Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	//Output:
	//Shaken not stirred
}
func BenchmarkCat(b *testing.B) {
	xs = strings.Split(c, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(c, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
