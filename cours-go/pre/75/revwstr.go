package main

import (
	"fmt"
	"os"
)

func plit(s, sep string) []string {
	res := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, str)
			i += len(sep) - 1
			str = ""
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
	return fin
}

func revwstr(s string) string {
	st := plit(s, " ")
	res := []string{}
	for i := len(st) - 1; i >= 0; i-- {
		res = append(res, st[i])
	}
	str := oin(res, " ")

	return str
}

func oin(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if i > 0 {
			str += sep + strs[i]
		} else {
			str += strs[i]
		}
	}
	return str
}

func main() {
	fmt.Println(plit("omar omar             pofnd okfjfe", " "))
	if len(os.Args) == 2 {
		fmt.Println(revwstr(os.Args[1]))
	}
}
