package main

import (
	"fmt"
	"os"
)

// Function to check if brackets are balanced
func isBalanced(s string) bool {
	stack := []rune{}
	bracketPairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char) // Push opening bracket
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false // Mismatched or unbalanced bracket
			}
			stack = stack[:len(stack)-1] // Pop from stack
		}
	}

	return len(stack) == 0 // Stack must be empty for a valid bracket sequence
}

func main() {
	if len(os.Args) < 2 {
		return // No arguments, print nothing
	}

	for _, arg := range os.Args[1:] {
		if isBalanced(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
