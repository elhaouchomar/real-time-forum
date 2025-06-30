package main

import (
    "fmt"
)

func FirstWord(s string) string {
    str := ""
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            str += string(s[i])
        } else if len(str) > 0 {
            break
        }
	}
    return str+"\n"
}

func main() {
    fmt.Print(FirstWord("   hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}