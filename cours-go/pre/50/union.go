package main

import (
	"fmt"
	"os"
)

func union(s1, s2 string) string{
	str := s1 + s2
	check := make(map[rune]bool)
	res := ""
	for _, c := range str {
		if !check[c] {
			res += string(c)
			check[c] = true
		}
	}
	return res
}

func main() {
	if len(os.Args) == 3 {
		fmt.Println(union(os.Args[1], os.Args[2]))
	}
}