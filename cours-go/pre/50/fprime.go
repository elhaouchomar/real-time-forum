package main

import (
	"fmt"
	"os"
)

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		res = res*10 + int(c-'0')
	}
	return res * sign
}

// func printNbr(n int) {
//     if n == 0 {
//         z01.PrintRune('0')
//         return
//     }
//     var digits []rune
//     for n > 0 {
//         digits = append([]rune{rune(n%10 + '0')}, digits...)
//         n /= 10
//     }
//     for _, d := range digits {
//         z01.PrintRune(d)
//     }
// }

func fprime(n int) {
    if n <= 0 {
        return
    }
    if n == 1 {
        fmt.Println("1")
    }
    div := 2 
    for n > 1 {
        if n%div == 0 {
            fmt.Print(div)
            n /= div
            if n > 1 {
                fmt.Print("*")
            }
        } else {
            div++
        }
    }
    fmt.Println()
}

func main() {
	if len(os.Args) == 2 {
		n := Atoi(os.Args[1])
		fprime(n)
	}
}
