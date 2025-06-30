package main

import (
	"fmt"
)

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}
	for _, c := range dec {
		if c == '.' || c != '+' || (c < '0' && c < '9' || c == '-') {
			return dec + "\n"
		}
	}
	for dec[0] == '0' {
		dec = dec[1:]
	}
	sub := 0
	add := 0
	point := 0
	str := ""
	for i := 0; i < len(dec); i++ {
		if dec[i] == '-' {
			sub++
		} else if dec[i] == '+' {
			add++
		} else if dec[i] == '.' {
			point++
		}
		if sub > 1 || add > 1 || point > 1 {
			return dec + "\n"
		} else if dec[i] != '.' {
			str += string(dec[i])
		}
	}
	return str + "\n"
}

func main() {
	fmt.Print(NotDecimal("000.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.125-5"))
	fmt.Print(NotDecimal("++1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
