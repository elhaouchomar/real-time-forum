package main

import (
	"fmt"
)


func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	if n < 0 {
		return "-" + Itoa(-n)
	}
	str := ""
	for n > 0 {
		str = string(n%10 + '0') + str
		n /=10
	}
	return str
}


func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}