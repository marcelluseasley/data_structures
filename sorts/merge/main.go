package main

import (
	"fmt"
)



// O (n log n)
func main() {
	//a := []int{56843,-76716,-42042,30846,50290,-27458,47259,91477,6775,59768,15476,-21788,-33386,42947,-27244,-45848,-7689,14298,-41118,-99082,-3641,-27944,-59233,25235,-44990,56287,-13288,8899,2579,-28954,27772,41360,34462,8174,53027,-22394,-91773,-42887,14202,30264,42811,-43774,-3273,11149,27266,-32695,31350,-5934,30010,-32840}
	a := []int{6,4,8,7,3,2,9,6,11,3,54,2,65,33,12,7,99,25}
	fmt.Println(MergeSort(a))
}

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice)/2

	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left []int, right []int) []int {

	size := len(left) + len(right)
	i, j := 0, 0

	slice := make([]int,size,size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}

	}
	return slice

}