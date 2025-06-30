package main

import (
	"fmt"
)

// ItoaBase converts an integer to a string representation in the given base (2 to 16)
func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	const digits = "0123456789ABCDEF"
	result := ""

	// Handle negative numbers
	if value < 0 {
		result = "-"
		value = -value
	}

	// Convert number to the given base
	var temp string
	for value > 0 {
		temp = string(digits[value%base]) + temp
		value /= base
	}

	// Handle zero case
	if temp == "" {
		temp = "0"
	}

	return result + temp
}

func main() {
	fmt.Println(ItoaBase(255, 2))  // Expected: "11111111"
	fmt.Println(ItoaBase(255, 16)) // Expected: "FF"
	fmt.Println(ItoaBase(42, 8))   // Expected: "52"
	fmt.Println(ItoaBase(-42, 10)) // Expected: "-42"
	fmt.Println(ItoaBase(0, 2))    // Expected: "0"
}
