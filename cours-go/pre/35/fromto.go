package main

import (
	"fmt"
)

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		return "-" + itoa(-n)
	}
	str := ""
	for n > 0 {
		str = string(n%10+'0') + str
		n /= 10
	}
	return str
}

func FromTo(from int, to int) string {
	str := ""
	for from > to {
		if from < 10 {
			str += "0" + itoa(from) + ", "
			from--
		} else {
			str += itoa(from) + ", "
			from--
		}
	}
	for to >= from {
		if from < 10 {
			str += "0" + itoa(from) + ", "
			from++
		} else {
			str += itoa(from) + ", "
			from++
		}
	}
	return str[:len(str)-2]
}

func main() {
	fmt.Println(FromTo(1, 10))
	fmt.Println(FromTo(10, 1))
	fmt.Println(FromTo(10, 10))
	fmt.Println(FromTo(100, 10))
}
