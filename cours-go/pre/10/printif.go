package main

import (
	"fmt"
)

func PrintIf(str string) string {
	if len(str) == 0 || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
