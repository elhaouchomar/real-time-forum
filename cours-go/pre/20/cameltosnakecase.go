package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string{
	str := ""
	for i := 0; i < len(s); i++ {
		 if i > 0 && (s[i] >= 'A' && s[i] <= 'Z'){
			str += "_"
		} else if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z'){
			return s
		}
		str += string(s[i])
	}
	return str
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}