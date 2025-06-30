package main

import (
	"fmt"
	"os"
)



func main() {
	if len(os.Args) != 2 {
		// إذا كان عدد المدخلات مختلف عن 1، يطبع Newline
		fmt.Println()
		return
	}

	// ناخدو السطر المدخل
	input := os.Args[1]

	// نفصل الكلمات
	words := Split(input, " ")

	// إذا كانو كلمات أكتر من واحدة، ندوّرها واحد لليسار
	if len(words) > 1 {
		// ناخدو الكلمات من العنصر الثاني لما الأخير
		rotatedWords := append(words[1:], words[0])

		// نطبعهم
		fmt.Println(Join(rotatedWords, " "))
	} else {
		// إذا كان غير كلمة واحدة، نطبعها كما هي
		fmt.Println(input)
	}
}
