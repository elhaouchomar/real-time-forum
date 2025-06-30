package main

import "fmt"

func fprime(n int) {
	if n <= 0  {
		return
	}
	if n == 1 {
		fmt.Println("1")
	}
	div := 2
	for n > 1 {
		if n % div == 0 {
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
	fmt.Println(17)
}