package main

import "fmt"

func BubbleSort(A []int, n int) []int {

	for k := 1; k < n-1; k++ {
		fmt.Println(A)
		swapped := false
		for i := 0; i < n-k-1; i++ {
			if A[i] > A[i+1] {
				A[i], A[i+1] = swap(A[i], A[i+1])
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return A

}

func swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func main() {

	a := []int{5, 7, 4, 3, 4, 6, 8, 2, 8, 4, 1, 5, 6}
	fmt.Println(BubbleSort(a, len(a)))

}
