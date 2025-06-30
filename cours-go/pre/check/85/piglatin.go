package main

import (
	"fmt"
	"os"
	"strings"
)

// Function to check if a character is a vowel
func isVowel(c rune) bool {
	return strings.ContainsRune("aeiouAEIOU", c)
}

// Function to convert a word to Pig Latin
func toPigLatin(word string) string {
	// Check if the word has at least one vowel
	hasVowel := false
	for _, c := range word {
		if isVowel(c) {
			hasVowel = true
			break
		}
	}
	if !hasVowel {
		return "No vowels"
	}

	// If the word starts with a vowel, add "ay" at the end
	if isVowel(rune(word[0])) {
		return word + "ay"
	}

	// Find the first vowel position
	for i, c := range word {
		if isVowel(c) {
			return word[i:] + word[:i] + "ay"
		}
	}

	return word // Should not reach here
}

func main() {
	// Ensure only one argument is passed
	if len(os.Args) != 2 {
		return
	}

	fmt.Println(toPigLatin(os.Args[1]))
}
