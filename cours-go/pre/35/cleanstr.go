package main

import (
	"fmt"
	"os"
)


func cleanStr(s string) string {
	str := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		} else if count > 0 && s[i] != ' ' && len(str) != 0 {
			str += " " + string(s[i])
			count = 0
		} else {
			str += string(s[i])
			count = 0
		}
	}
	return str
}

func main() {
	if len(os.Args) !=  2 {
		return
	}
	arg :=  os.Args[1] 
	res := cleanStr(arg)
	fmt.Println(res)
}