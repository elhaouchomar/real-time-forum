package main

import (
	"fmt"
	"os"
)

func hiddenp(s1, s2 string) {
	if len(s1) == 0 {
		fmt.Println(1)
	}
	j := 0

	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[j] {
			j++
		}
		if j == len(s1) {
			fmt.Println(1)
			return
		}
	}
	fmt.Println(0)
}

func main() {
	if len(os.Args) == 3 {
		hiddenp(os.Args[1], os.Args[2])
	}
}
