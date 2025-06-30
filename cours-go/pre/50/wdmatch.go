package main

import (
	"fmt"
	"os"
)

func wdmatch(s1, s2 string) {
	if len(s1) == 0 {
		fmt.Println("")
		return
	}
	
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}
	
	if j == len(s1) {
		fmt.Println(s1)
	}
}

func main() {
	if len(os.Args) == 3 {
		wdmatch(os.Args[1], os.Args[2])
	}
}