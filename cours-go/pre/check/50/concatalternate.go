package main

import "fmt"

// ConcatAlternate fusionne deux slices en alternant les éléments
func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	len1, len2 := len(slice1), len(slice2)

	// Si slice1 est plus grand ou si les deux ont la même taille, on commence par slice1
	if len1 >= len2 {
		for i := 0; i < len1 || i < len2; i++ {
			if i < len1 {
				result = append(result, slice1[i])
			}
			if i < len2 {
				result = append(result, slice2[i])
			}
		}
	} else {
		// Sinon, on commence par slice2
		for i := 0; i < len2 || i < len1; i++ {
			if i < len2 {
				result = append(result, slice2[i])
			}
			if i < len1 {
				result = append(result, slice1[i])
			}
		}
	}

	return result
}

func main() {
	// Cas de test 1
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))

	// Cas de test 2
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))

	// Cas de test 3
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))

	// Cas de test 4
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
