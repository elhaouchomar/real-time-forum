package main

import (
	"fmt"
	"os"
)

func toLow(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		c += 32
	}
	return c
}

func toCap(c rune) rune {
	if c >= 'a' && c <= 'z' {
		c -= 32
	}
	return c
}

func reversestrcap(s string) string {
	if len(s) == 0 {
		return "\n"
	}
	str := ""
	for _, c := range s {
		str += string(toLow(c))
	}
	res := ""

	for i := 0; i < len(str)-1; i++ {
		if str[i] >= 'a' && str[i] <= 'z' && str[i+1] == ' ' {
			res += string(toCap(rune(str[i])))
		} else {

			res += string(str[i])
		}
	}

	// if len(str) > 0 {
	// 	if str[len(str)-1] >= 'a' && str[len(str)-1] <= 'z' {
	// 		res += string(toCap(rune(str[len(str)-1])))
	// 	} else {

	// 		res += string(str[len(str)-1])
	// 	}
	// }
	return res
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(reversestrcap(os.Args[i]))
	}
}
