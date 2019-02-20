package main

import (
	"fmt"
	"sort"
)

func main() {
	si := []int{34, 82, 99, 12, 56, 1, 78, 32, 10, 02}
	ss := []string{"hello", "paresh", "columbas", "tree", "junior", "border", "america", "zebra"}

	fmt.Println("Before Sorting", si)
	sort.Ints(si)
	fmt.Println("After Sorting", si)

	fmt.Println("Before Sorting", ss)
	sort.Strings(ss)
	fmt.Println("After Sorting", ss)

}
