package main

import "testing"

// testing mysum method of testing execericse
/*func TestMySum(t *testing.T) {
	x := mysum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, "Got ", x)
	}
}*/

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{2, 7}, 9},
		test{[]int{-1, 0, 21}, 20},
		test{[]int{1, 2}, 3},
	}

	for _, v := range tests {
		x := mysum(v.data...)
		if x != v.answer {
			t.Error("Expected ", v.answer, "Got ", x)
		}
	}
}
