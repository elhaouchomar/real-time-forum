package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
    count := 0
	for _, ch := range str {
		if ch == c {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
