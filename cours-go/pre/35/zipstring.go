package main

import (
	"fmt"
)

func itoaa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		return "-" + itoaa(-n)
	}
	str := ""
	for n > 0 {
		str = string(n%10+'0') + str
		n /= 10
	}
	return str
}

func ZipString(s string) string {
	count := 1
	str := ""
	for i := 1; i < len(s); i++ {
		if (s[i] >= 'a' &&  s[i] <= 'z') || (s[i] >= 'A' &&  s[i] <= 'Z') {
			if s[i] == s[i-1] {
				count++
			} else {
				str += itoaa(count) + string(s[i-1])
				count = 1
			}
		} else {
			str += itoaa(count) + string(s[i])
		}
	}
	str  += itoaa(count) + string(s[len(s)-1])
	return str
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}