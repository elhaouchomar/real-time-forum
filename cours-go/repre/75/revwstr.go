package main

import (
	"fmt"
	"os"
)

func split(s, sep string) []string {
	res := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, str)
			str = ""
			i += len(sep) - 1
		} else {

			str += string(s[i])
		}
	}
	res = append(res, str)
	fin := []string{}
	for i := 0; i < len(res); i++ {
		if res[i] != "" {
			fin = append(fin, res[i])
		}
	}
	return res
}

func revstr(s string) string {
	str := split(s, " ")
	ss := ""
	for i := len(str) - 1; i >= 0; i-- {
		if i != len(str) -1 {
			ss += " " + str[i]
		} else {
			ss +=  str[i]
		}
	}
	return ss
}
func main() {
	if len(os.Args) == 2 {
		fmt.Println(revstr(os.Args[1]))
	}
}
