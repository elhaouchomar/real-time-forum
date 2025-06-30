package main

import (
	"fmt"
)

func count(s string) int {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			count = int(c) - int('a') + 1
		} else if c >= 'A' && c <= 'Z' {
			count = int(c) - int('A') + 1
		}
	}
	return count
}

func RepeatAlpha(s string) string {
	str := ""
	for _, c := range s {
		count := count(string(c))
		for count > 0 {
			str += string(c)
			count--
		}
	}
	return str
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
