package main

import "fmt"

func main() {
	arr := [8]int64{23, 05, 12, 64, 89, 21, 33, 99}
	fmt.Println("Sorting the array by selection sort algorithm")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for j := 0; j < len(arr)-1; j++ {
		min := j
		///fmt.Printf("the value of j is %d\n", arr[j])
		//This inner for loop is to get only the min value from the array
		for k := j + 1; k < len(arr); k++ {
			//fmt.Printf("the value of k is %d %d\n", arr[k], arr[min])
			if arr[k] < arr[min] {
				min = k
			}
		}
		//fmt.Printf("the value of k is %d\n", arr[k])
		temp := arr[min]
		arr[min] = arr[j]
		arr[j] = temp

	}

	fmt.Println("Sorted array by selection sort algorithm")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
