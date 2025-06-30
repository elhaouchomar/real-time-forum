package main

import "fmt"

// CanJump détermine si on peut atteindre la dernière position de l'array
func CanJump(nums []uint) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	pos := 0 // Position actuelle
	for pos < n-1 {
		if nums[pos] == 0 {
			return false // Bloqué, impossible d'avancer
		}

		pos += int(nums[pos]) // Avancer de `nums[pos]` positions

		if pos >= n-1 {
			return true // Atteint ou dépassé la dernière position
		}
	}

	return false
}

func main() {
	// Cas de test 1
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // Devrait afficher: true

	// Cas de test 2
	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // Devrait afficher: false

	// Cas de test 3
	input3 := []uint{0}
	fmt.Println(CanJump(input3)) // Devrait afficher: true
}
