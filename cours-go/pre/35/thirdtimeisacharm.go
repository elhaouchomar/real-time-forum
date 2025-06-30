package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	s := ""
	count := 0
	for i := 0; i < len(str); i++ {
		count++
		if count % 3 == 0 {
			s += string(str[i])
		}
	}
	return s+"\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}