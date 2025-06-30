package main

import (
	"fmt"
	"strings"
)

func clean(s string) string {
	for s != "" && s[0] == ' ' {
		s = s[1:]
	}
	for s != "" && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func oneSpace(s string) string {
	res := ""
	check := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			check++
		}
		if s[i] != ' ' {
			res += string(s[i])
			check = 0
		}
		if check > 0 {
			res += " "
		}
	}
	return res
}

func WordFlip(str string) string {
	str = clean(str)
	fmt.Println("clean : ", str)
	str = oneSpace(str)
	fmt.Println("onespace : ", str)
	s := strings.Split(str, " ")
	fmt.Println("split : ", s, len(s))
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] + " "

	}
	return res[:len(res)-1] + "\n"
}

func split(s, sep string) []string {
	res := []string{}
	st := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, st)
			st = ""
			i += len(s) -1
		} else {
			st += string(s[i])
		}
	}
	res = append(res, st)
	return res
}

func join(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if i > 0 {
			str += sep + strs[i]
		}else {
			str += strs[i]
		}
	}
	return str
}

func main() {
	expected := []string{"apple", "banana", "cherry", "date"}
	fmt.Println(join(expected, "       "))
	fmt.Println(split("First second last", " "))
	fmt.Print(WordFlip("First second last"))
	// fmt.Print(WordFlip(""))
	// fmt.Print(WordFlip("     "))
	// fmt.Print(WordFlip(" hello  all  of  you! "))
}
