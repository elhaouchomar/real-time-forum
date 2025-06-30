package main

import (
	"fmt"
)

func WordFlip(str string) string {

	if str == "" {
		return "Invalid Output"
	}
	words := Split(str, ' ')

	reversed := ""
	for i := len(words) - 1; i >= 0; i-- {
		reversed += words[i] + " "
	}
	if reversed[len(reversed)-1] == ' ' && len(reversed) > 1{
		return reversed[:len(reversed)-1]
	}
	return reversed + "\n"
}

func Split(str string, delimiter rune) []string {
	var words []string
	var currentWord []rune

	// Iterate over the string and build words
	for _, char := range str {
		if char == delimiter {
			// If we encounter the delimiter, add the current word to the result slice
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		} else {
			// Otherwise, add the character to the current word
			currentWord = append(currentWord, char)
		}
	}

	// Add the last word to the result if any
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}

func Join(words []string, delimiter string) string {
	var result string

	// Iterate over the words and concatenate them with the delimiter
	for i, word := range words {
		if i != 0 {
			// Add the delimiter between words, but not before the first word
			result += delimiter
		}
		result += word
	}

	return result
}

func main() {

	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
