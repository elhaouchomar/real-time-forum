package main

import (
	"fmt"
)

func LastWord(s string) string{
	str := ""
	for i := len(s) -1; i >= 0; i-- {
		if s[i] != ' ' {
			str += string(s[i])
		} else if len(str) > 0 {
			break
		}
	}
	res := ""
	for i := len(str)-1; i >= 0; i-- {
		res += string(str[i])
	}
	return res+"\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}