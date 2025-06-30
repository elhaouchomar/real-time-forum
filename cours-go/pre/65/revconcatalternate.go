package main

import (
	"fmt"
)

func RevConcatAlternate(slice1, slice2 []int) []int {
	l1, l2 := len(slice1), len(slice2)
	res := []int{}
	if l2 > l1 {
		for i := l2-1; i >= 0; i-- {
			if i < l1{
				res = append(res, slice1[i])
			}
			if i < l2 {
				res = append(res, slice2[i])
			}
		}

	} else {
		for i := l1-1; i >= 0; i-- {
			if i < l2 {
				res = append(res, slice2[i])
			}
			if i < l1 {
				res = append(res, slice1[i])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
