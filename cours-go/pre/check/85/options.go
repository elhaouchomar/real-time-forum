package main

import (
	"fmt"
	"os"
)

const validOptions = "abcdefghijklmnopqrstuvwxyz"

func printOptions() {
	fmt.Println("options: " + validOptions)
}

func printBinary(n int) {
	fmt.Printf("%08b %08b %08b %08b\n", (n>>24)&0xFF, (n>>16)&0xFF, (n>>8)&0xFF, n&0xFF)
}

func isValidOption(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func main() {
	if len(os.Args) == 1 {
		printOptions()
		return
	}

	options := 0

	for _, arg := range os.Args[1:] {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		// If the "-h" flag appears anywhere, print options and exit
		if arg == "-h" || (len(arg) > 2 && arg[1] == 'h') {
			printOptions()
			return
		}

		// Process each character after "-"
		for _, char := range arg[1:] {
			if !isValidOption(char) {
				fmt.Println("Invalid Option")
				return
			}
			options |= (1 << (char - 'a'))
		}
	}

	printBinary(options)
}
