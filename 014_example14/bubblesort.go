package main

import "fmt"

//Bubble Sort :  is the simpest algorithm that works on repeated swipping the adjecent elements if they are in the wrong order.
func main() {
	/*for i := 0; i <= 10; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf(" the outloop value is %d and inner loop value is %d\n", i, j)
		}
	}
	the outloop value is 0 and inner loop value is 0
	the outloop value is 0 and inner loop value is 1
	the outloop value is 0 and inner loop value is 2
	the outloop value is 0 and inner loop value is 3
	the outloop value is 1 and inner loop value is 0
	the outloop value is 1 and inner loop value is 1
	the outloop value is 1 and inner loop value is 2
	the outloop value is 1 and inner loop value is 3
	the outloop value is 2 and inner loop value is 0
	the outloop value is 2 and inner loop value is 1
	the outloop value is 2 and inner loop value is 2
	the outloop value is 2 and inner loop value is 3
	the outloop value is 3 and inner loop value is 0
	the outloop value is 3 and inner loop value is 1
	the outloop value is 3 and inner loop value is 2
	the outloop value is 3 and inner loop value is 3
	the outloop value is 4 and inner loop value is 0
	the outloop value is 4 and inner loop value is 1
	the outloop value is 4 and inner loop value is 2
	the outloop value is 4 and inner loop value is 3
	the outloop value is 5 and inner loop value is 0
	the outloop value is 5 and inner loop value is 1
	the outloop value is 5 and inner loop value is 2
	the outloop value is 5 and inner loop value is 3
	the outloop value is 6 and inner loop value is 0
	the outloop value is 6 and inner loop value is 1
	the outloop value is 6 and inner loop value is 2
	the outloop value is 6 and inner loop value is 3
	the outloop value is 7 and inner loop value is 0
	the outloop value is 7 and inner loop value is 1
	the outloop value is 7 and inner loop value is 2
	the outloop value is 7 and inner loop value is 3
	the outloop value is 8 and inner loop value is 0
	the outloop value is 8 and inner loop value is 1
	the outloop value is 8 and inner loop value is 2
	the outloop value is 8 and inner loop value is 3
	the outloop value is 9 and inner loop value is 0
	the outloop value is 9 and inner loop value is 1
	the outloop value is 9 and inner loop value is 2
	the outloop value is 9 and inner loop value is 3
	the outloop value is 10 and inner loop value is 0
	the outloop value is 10 and inner loop value is 1
	the outloop value is 10 and inner loop value is 2
	the outloop value is 10 and inner loop value is 3
	*/
	arr := [7]int64{88, 30, 15, 90, 21, 9, 33}
	fmt.Println("Sorting array for bubble sort")
	for k := 0; k < len(arr); k++ {
		fmt.Println(arr[k], "\t")
	}
	for i := 0; i < len(arr)-1; i++ {
		//fmt.Printf("the value of i is %d\t", arr[i])
		for j := 0; j < len(arr)-i-1; j++ {
			//fmt.Printf("The value of j is %d\n", arr[j])
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println("Sorted array for bubble sort")
	for k := 0; k < len(arr); k++ {
		fmt.Println(arr[k], "\t")
	}

}

/*
Sorting array for bubble sort
88
30
15
90
21
9
33
Sorted array for bubble sort
9
15
21
30
33
88
90
*/
