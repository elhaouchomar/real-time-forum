package main

import (
	"fmt"
	"os"
)



func spl(s, sep string) []string {
	res := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, str)
			str = ""
			i += len(sep) -1
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
func rosting(s string) string {
	str := spl(s, " ")
	st := ""
	for i := 0; i < len(str); i++ {
		if i > 0 {
			st += " " + str[i]
		} else {
			st += str[i]
		}
	}
	return st
}

func main() {
	fmt.Println(spl("jdbbjd bdf                                         fbd jehfj", " "))
	// fmt.Println()
	if len(os.Args) == 2 {
		fmt.Println(rosting(os.Args[1]))
	}
	
}