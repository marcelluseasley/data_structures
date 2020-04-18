// Quick Sort in Golang
package main

import (
	"fmt"
)



func main() {
	A := []int{7, 2, 1, 6, 8, 5, 3, 4}
	QuickSort(A, 0, len(A)-1)
	fmt.Println(A)

}

func QuickSort(A []int, start, end int) {

	if start < end {
		pIndex := Partition(A, start, end)

		QuickSort(A, start, pIndex-1)
		QuickSort(A, pIndex+1, end)
	}
}

func Partition(A []int, start, end int) int {
	pivot := A[end]
	pIndex := start

	for i := start; i < end; i++ {
		if A[i] <= pivot {
			A[i], A[pIndex] = swap(A[i], A[pIndex])
			pIndex++
		}
	}
	A[pIndex], A[end] = swap(A[pIndex], A[end])

	return pIndex
}

func swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}