package main

import (
	"fmt"
)

func RevConcatAlternate(slice1, slice2 []int) []int {
	// نحدد الطول ديال كل لائحة
	len1, len2 := len(slice1), len(slice2)
	var result []int

	// نخدم على اللائحة اللي عندها طول أكبر
	if len1 >= len2 {
		slice1, slice2 = reverse(slice1), reverse(slice2)
	} else {
		slice2, slice1 = reverse(slice2), reverse(slice1)
	}

	i, j := 0, 0

	// ندير alternation بين القيم ديال اللائحتين
	for i < len(slice1) && j < len(slice2) {
		result = append(result, slice1[i])
		i++
		if j < len(slice2) {
			result = append(result, slice2[j])
			j++
		}
	}

	// نضيف القيم المتبقية من اللائحة الأطول
	for i < len(slice1) {
		result = append(result, slice1[i])
		i++
	}

	return result
}

// دالة لعكس اللائحة
func reverse(slice []int) []int {
	n := len(slice)
	reversed := make([]int, n)
	for i := 0; i < n; i++ {
		reversed[i] = slice[n-1-i]
	}
	return reversed
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))          // [3 6 2 5 1 4]
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9})) // [9 8 7 3 6 2 5 1 4]
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))       // [8 9 3 2 5 1 4]
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))                 // [3 2 1]
}
