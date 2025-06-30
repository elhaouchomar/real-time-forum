package main

import (
	"fmt"
	"os"
)

func interpretBrainfuck(code string) {
	const memorySize = 2048
	memory := make([]byte, memorySize) 
	ptr := 0 // Memory pointer
	loopStack := []int{} // Stack for loops
	
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr = (ptr + 1) % memorySize
		case '<':
			ptr = (ptr - 1 + memorySize) % memorySize
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Printf("%c", memory[ptr])
		case ',':
			// Read a byte from stdin
			var input [1]byte
			os.Stdin.Read(input[:])
			memory[ptr] = input[0]
		case '[':
			if memory[ptr] == 0 {
				// Skip to the matching closing bracket
				depth := 1
				for depth > 0 {
					i++
					if i >= len(code) {
						return // End of code, exit
					}
					if code[i] == '[' {
						depth++
					} else if code[i] == ']' {
						depth--
					}
				}
			} else {
				// Push the current position onto the stack
				loopStack = append(loopStack, i)
			}
		case ']':
			if memory[ptr] != 0 {
				// Jump back to the matching opening bracket
				if len(loopStack) > 0 {
					i = loopStack[len(loopStack)-1]
				}
			} else {
				// Pop the stack when exiting the loop
				if len(loopStack) > 0 {
					loopStack = loopStack[:len(loopStack)-1]
				}
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [brainfuck code]")
		return
	}
	interpretBrainfuck(os.Args[1])
}