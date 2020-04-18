package main

import (
	"fmt"
)

func SelectionSort(A []int, n int) []int {


	for i := 0; i < n-1; i++ {

		iMin := i
		for j := i+1; j < n; j++ {
			if A[j] < A[iMin] {
				iMin = j
			}
		}
		temp := A[i]
		A[i] = A[iMin]
		A[iMin] = temp
	}
	return A
}

func main() {

	a := []int{5,7,4,3,4,6,8,2,8,4,1,5,6}
	fmt.Println(SelectionSort(a, len(a)))
}