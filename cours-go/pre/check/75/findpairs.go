package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// دالة لاستخراج الأعداد من النص
func parseArray(input string) ([]int, error) {
	input = strings.TrimSpace(input)
	if len(input) < 2 || input[0] != '[' || input[len(input)-1] != ']' {
		return nil, fmt.Errorf("Invalid input.")
	}

	numStrs := strings.Split(input[1:len(input)-1], ",")
	var numbers []int

	for _, numStr := range numStrs {
		numStr = strings.TrimSpace(numStr)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("Invalid number: %s", numStr)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// دالة لاستخراج target sum
func parseTarget(input string) (int, error) {
	input = strings.TrimSpace(input)
	target, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Invalid target sum.")
	}
	return target, nil
}

// دالة لإيجاد الأزواج
func findPairs(arr []int, targetSum int) [][]int {
	var pairs [][]int
	used := make([]bool, len(arr))

	for i := 0; i < len(arr); i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if used[j] {
				continue
			}
			if arr[i]+arr[j] == targetSum {
				pairs = append(pairs, []int{i, j})
				used[i], used[j] = true, true
				break
			}
		}
	}
	return pairs
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arr, err := parseArray(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	targetSum, err := parseTarget(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	pairs := findPairs(arr, targetSum)
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", targetSum, pairs)
	}
}
