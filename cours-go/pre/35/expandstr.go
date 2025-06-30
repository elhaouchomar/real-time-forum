package main

import (
	"os"

	"github.com/01-edu/z01"
)

func expandStr(s string) string {
	count := 0
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		} else if count > 0 && s[i] != ' ' && len(str) > 0 {
			str += "    " + string(s[i])
			count = 0
		} else {
			str += string(s[i])
			count = 0
		}
	}
	return str+"\n"
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := expandStr(os.Args[1])

	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}

}
