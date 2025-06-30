package main

import "fmt"

func contains(s string, c rune) bool {
	for _, ch := range s {
		if c == ch {
			return true
		}
	}
	return false
}

func inter(s1, s2 string) {
	seen := make(map[rune]bool)
	for _, c := range s1 {
		if !seen[c] && contains(s2, c) {
			fmt.Print(c)
			seen[c] = true
		}
	}
	fmt.Println()
}

func main() {
	
}