package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}
	charMap := make(map[rune]int)
	for _, ch := range str1 {
		charMap[ch]++
	}

	for _, ch := range str2 {
		charMap[ch]++
	}

	uni := 0

	for _, count := range charMap {
		if count == 1 {
			uni++
		}
	}
	return uni
}
 
func main() {
	fmt.Println(WeAreUnique("foo", "boo")) 
	fmt.Println(WeAreUnique("", ""))       
	fmt.Println(WeAreUnique("abc", "def")) 
}
