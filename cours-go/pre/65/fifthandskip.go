package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	s := ""
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			s += string(str[i])
		}
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if (i+1) % 6 != 0{
			res += string(s[i])
			} else if len(s) -1 > i{
			res += " "

		}
	}
	return res+"\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("123456"))
}