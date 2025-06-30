package main

import (
	"fmt"
)

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	} else if n%2 == 0 && n%3 == 0 {
		return "fish and chips\n"
	} else if n%2 == 0 {
		return "fish"
	} else if n%3 == 0 {
		return "chips"
	} else {
		return "error: non divisible\n"
	}
}

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
