package main

import (
	"os"

	"github.com/01-edu/z01"
)

func contains(s string, c rune) bool {
	for _, ch := range s {
		if ch == c {
			return true
		}
	}
	return false
}

func inter(s1, s2 string) {
	seen := make(map[rune]bool)
	for _, c := range s1 {
		if !seen[c] && contains(s2, c) {
			z01.PrintRune(c)
			seen[c] = true
		}
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 3 {
		inter(os.Args[1], os.Args[2])
	}
}
