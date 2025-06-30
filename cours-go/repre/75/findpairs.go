package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	res := 0
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
			s = s[1:]
		} else {
			s = s[1:]
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		res = res*10 + int(s[i]-'0')
	}
	return res * sign
}

func fields(s, sep string) []string {
	res := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			res = append(res, str)
			str = ""
			i += len(sep) - 1
		} else {
			str += string(s[i])
		}
	}
	res = append(res, str)
	fin := []string{}
	for i := 0; i < len(res); i++ {
		if res[i] != "" {
			fin = append(fin, res[i])
		}
	}
	return fin
}

func join(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if i != 0 {
			str += sep + strs[i]
		} else {
			str += strs[i]
		}
	}
	return str
}

func main() {
	s, sep := "ofbokfb dvkdvkdm dvodvd vdovjdobvdob      dkvdkv ovdov", " "
	sp := fields(s, sep)
	fmt.Println(sp)
	fmt.Println(join(sp, " "))
	fmt.Println(atoi("-545445"))
	// if len(os.Args) != 3 {
	// 	fmt.Println("Invalid input.")
	// }
	f := os.Args[1]
	se := os.Args[2]
	if len(se) > 1 && atoi(se) == 0 {
		fmt.Println("Invalid input.")
		return
	}
	if f[0] != '[' && f[len(f)-1] != ']' {
		fmt.Println("Invalid input.")
	} else {
		f = f[1:len(f)-1]
		f := fields(f, ", ")
		fmt.Println(f)
		res := []int{}
		for i := 0; i < len(f); i++ {
			res = append(res, atoi(string(f[i])))
		}
		fmt.Println(res)
		fin := []int{}
		rw := [][]int{}
		for i := 0; i < len(res)-1 ; i++ {
			for j := i+1 ; j < len(res); j++ {
				if res[i] + res[j] == atoi(se) {
					fin = append(fin, i, j)
					rw = append(rw, fin)
					fin = []int{}
				}
 			}
			
		}
		if len(rw) == 0 {
			fmt.Println("No pairs found.")
			return
		}
 		fmt.Printf("Pairs with sum %v: %v\n", se, rw)
	}

}

