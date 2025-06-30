package main

import "fmt"

// ConcatSlice concatène deux slices en un seul
func ConcatSlice(slice1, slice2 []int) []int {
	// Concaténation de slice1 et slice2
	return append(slice1, slice2...)
}

func main() {
	// Cas de test 1
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))

	// Cas de test 2
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))

	// Cas de test 3
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
