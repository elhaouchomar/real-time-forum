package main

import (
	"fmt"
)

// Splits the string by the given delimiter while handling multiple spaces
func Split(str string, delimiter rune) []string {
	var words []string
	var currentWord []rune

	for _, char := range str {
		if char == delimiter {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		} else {
			currentWord = append(currentWord, char)
		}
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}

// Joins the words with the given delimiter
func Join(words []string, delimiter string) string {
	if len(words) == 0 {
		return ""
	}

	result := words[0]
	for i := 1; i < len(words); i++ {
		result += delimiter + words[i]
	}
	return result
}

// Flips the words in a string
func WordFlip(s string) string {
	words := Split(s, ' ') // Split into words
	if len(words) == 0 {
		return "Invalid Output\n"
	}

	// Reverse the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return Join(words, " ") + "\n"
}

func main() {
	fmt.Print(WordFlip("First second last"))      // "last second First\n"
	fmt.Print(WordFlip(""))                       // "Invalid Output\n"
	fmt.Print(WordFlip("     "))                  // "Invalid Output\n"
	fmt.Print(WordFlip(" hello  all  of  you! ")) // "you! of all hello\n"
}
