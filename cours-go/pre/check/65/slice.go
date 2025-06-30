package main

import (
	"fmt"
)

// الدالة Slice
func Slice(a []string, nbrs ...int) []string {
	n := len(a)

	// التحقق من عدد الأرقام المدخلة
	if len(nbrs) == 0 || len(nbrs) > 2 {
		return nil
	}

	// استخراج البداية والنهاية
	start, end := nbrs[0], n
	if len(nbrs) == 2 {
		end = nbrs[1]
	}

	// التعامل مع القيم السالبة
	if start < 0 {
		start += n
	}
	if end < 0 {
		end += n
	}

	// التحقق من صحة القيم
	if start < 0 || start >= n || end < 0 || end > n || start >= end {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))        // ["algorithm", "ascii", "package", "golang"]
	fmt.Printf("%#v\n", Slice(a, 2, 4))     // ["ascii", "package"]
	fmt.Printf("%#v\n", Slice(a, -3))       // ["ascii", "package", "golang"]
	fmt.Printf("%#v\n", Slice(a, -2, -1))   // ["package"]
	fmt.Printf("%#v\n", Slice(a, 2, 0))     // nil
}
