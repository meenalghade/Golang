//Package word provide custom functions
package word

import "strings"

//This is the method used to get the string count
func UserCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns number of words in string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
