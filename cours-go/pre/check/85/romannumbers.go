package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Roman numeral mapping
var romanMap = []struct {
	value int
	symbol string
}{
	{1000, "M"}, {900, "(M-C)"}, {500, "D"}, {400, "(D-C)"},
	{100, "C"}, {90, "(C-X)"}, {50, "L"}, {40, "(L-X)"},
	{10, "X"}, {9, "(X-I)"}, {5, "V"}, {4, "(V-I)"},
	{1, "I"},
}

// Convert integer to Roman numeral with breakdown
func intToRoman(num int) (string, string) {
	var roman strings.Builder
	var breakdown strings.Builder

	for _, entry := range romanMap {
		for num >= entry.value {
			if breakdown.Len() > 0 {
				breakdown.WriteString("+")
			}
			breakdown.WriteString(entry.symbol)
			roman.WriteString(strings.ReplaceAll(entry.symbol, "(", ""))
			roman.WriteString(strings.ReplaceAll(entry.symbol, ")", ""))
			num -= entry.value
		}
	}

	return breakdown.String(), roman.String()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	breakdown, roman := intToRoman(num)
	fmt.Println(breakdown)
	fmt.Println(roman)
}
