package main

import (
	"fmt"
)

func CanJump(steps []uint) bool {
	n := len(steps)

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	pos := 0 
	for pos <= n {
		if steps[pos] == 0 {
			return false
		}
		pos += int(steps[pos])
		if pos >= n {
			return true
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}