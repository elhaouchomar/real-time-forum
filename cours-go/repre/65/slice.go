package main

import (
	"fmt"
)

func Slice(a []string, nbrs ...int) []string {
	n1 := nbrs[0]
	n2 := nbrs[1]
	if len(nbrs) == 1 && n1 >= 0 {
		return a[n1:]
	}else if len(nbrs) > 1 && n1 >= 0 && n2 >= 0 {
		if n2 < n1 {
			return nil
		}
		return a[n1:n2]
	} else if len(nbrs) > 1 && n1 < 0 && n2 < 0 {
		n1 += len(a)
		n2 += len(a)
		return a[n1:n2]
	} else if len(nbrs) == 1 && n1 < 0 {
		n1 += len(a)
		return a[n1:]
	}
	return a
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
