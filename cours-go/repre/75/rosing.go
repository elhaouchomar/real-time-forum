package main

import (
	"fmt"
	"os"
)

func split(s, sep string) string {
	res := []string{}
	st := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, st)
			st = ""
			i += len(sep) - 1
		} else {
			st += string(s[i])
		}
	}
	res = append(res, st)
	fin := []string{}
	for i := 0; i < len(res); i++ {
		if res[i] != "" {
			fin = append(fin, res[i])
		}
	}
	stt := ""
	for i := 1; i < len(fin); i++ {
		if i != 0 {
			stt += " " + fin[i]
		} else {
			stt += fin[i]
		}
	}
	
	stt += " " + fin[0]
	for stt[0] == ' ' {
		stt = stt[1:]
	}
	return stt
}

// func rosting(s string) string {
// 	res := split(s, " ")

// }

func main() {
	// fmt.Println(split("dfdfkjd dfhdshf      jfd", " "))
	if len(os.Args) == 2 {
		str := split(os.Args[1], " ")
		fmt.Println(str)
		fmt.Println(len(str))
	}
}
