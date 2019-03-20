// "Package  acdc ready to be used"
package acdc

// sum added unlimited addition of numbers
func sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
