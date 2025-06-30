package main

import (
	"fmt"
	"unicode"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	hasDot := false
	result := ""

	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			hasDot = true
			continue
		}
		if !unicode.IsDigit(rune(dec[i])) && dec[i] != '-' {
			return dec + "\n"
		}
		result += string(dec[i])
	}

	if !hasDot || result[len(result)-1] == '0' {
		return dec + "\n"
	}

	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
