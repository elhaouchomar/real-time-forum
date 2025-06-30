package main

import (
	"fmt"
)

func HashCode(dec string) string {
	str := ""
	for i := 0; i < len(dec); i++ {
		if dec[i] >= 32 && dec[i] <= 126 {
			str += string((int(dec[i]) + len(dec))%127)
		} else {
			str += string(33)
		}
	}
	return str
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}