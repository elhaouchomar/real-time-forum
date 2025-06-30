package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(s1, s2 string) {
	if len(s1) < 2 {
		fmt.Println("Invalid input.\n")
		return
	}
	n, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	arr := []string{}
	str := ""
	for i := 0; i < len(s1); i++ {
		if s1[0] != '[' && s1[len(s1)-1] != ']' {
			fmt.Println("Invalid input.")
			return
		} else if i > 0 && i < len(s1)-1 {
			str += string(s1[i])
		}
	}
	arr = strings.Split(str, ", ")
	res := []int{}
	for i := 0; i < len(arr); i++ {
		nb, er := strconv.Atoi(arr[i])
		if er != nil {
			fmt.Printf("Invalid number: %v\n", strings.TrimSpace(arr[i]))
			return
		}
		s := ""
		for s != "" && s[0] == ' ' {
			s = s[1:]
		}
		res = append(res, nb)
	}
	fmt.Println(res)
	check := false
	array := []int{}
	fin := [][]int{}
	for i := 0; i < len(res)-1; i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i]+res[j] == n {
				array = append(array, i, j)
				fin = append(fin, array)
				array = []int{}
				check = true
			}
		}
	}
	if !check {
		fmt.Println("No pairs found.")
		return
	}
	fmt.Printf("Pairs with sum %d: %v", n, fin)
	fmt.Println()
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}
	parseInput(os.Args[1], os.Args[2])
}
