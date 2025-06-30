package main

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(base[int(arr[i]/16)]))
		z01.PrintRune(rune(base[int(arr[i]%16)]))
		if !(i == 3 || i == 7 || i == 9) {
			z01.PrintRune(' ')
		} else if i == 3 || i == 7 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')

	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')

}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
