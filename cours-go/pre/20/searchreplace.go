package main

import (
	"os"
 "github.com/01-edu/z01"

)

func searchreplace(s, old, new string) string{
	str := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == old {
			str += new
		} else {
			str += string(s[i])
		}
	}
	return str
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	arg := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]
	res := searchreplace(arg, old, new)
	for i := 0; i < len(res); i++ {
		z01.PrintRune(rune(res[i]))
	}
	z01.PrintRune('\n')
}