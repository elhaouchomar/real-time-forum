package main

import (
	"fmt"
)

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) > 1 && nbrs[0] >= 0 && nbrs[1] >= 0 {
		if nbrs[1] < nbrs[0] {
			return nil
		}
		return a[nbrs[0]:nbrs[1]]
	} else if len(nbrs) == 1 && nbrs[0] > 0 {
		return a[nbrs[0]:]
	} else if len(nbrs) > 1 && nbrs[0] < 0 && nbrs[1] < 0 {
		nbrs[0] = nbrs[0] + len(a)
		nbrs[1] = nbrs[1] + len(a)
		return a[nbrs[0]:nbrs[1]]
	} else if len(nbrs) == 1 && nbrs[0] < 0 {
		nbrs[0] = nbrs[0] + len(a)
		return a[nbrs[0]:]
	}
	return a
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 2))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 3))
}
