package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		// إذا كان عدد المدخلات مختلف عن 1، ما كيطبعش حاجة
		return
	}

	// ناخدو السطر المدخل
	input := os.Args[1]

	// نفصل الكلمات
	words := strings.Fields(input)

	// نطبع الكلمات من آخر وحدة للأولى
	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words)-1 {
			fmt.Print(" ")
		}
		fmt.Print(words[i])
	}

	// نطبع newline في الأخير
	fmt.Println()
}
