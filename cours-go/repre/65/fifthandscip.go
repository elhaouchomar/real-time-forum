package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	s := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			s += string(str[i])
		}
	}
	n := ""
	for i := 0; i < len(s); i++ {
		if(i+1) % 6 != 0 {
			n += string(s[i])
		} else if len(s) > i {
			n += " "
		}
	
	}
	return n+"\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}