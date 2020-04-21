// 2020/3/8
package main

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	ans := ""
	if n%2 == 1 {
		ans = strings.Repeat("a", n)
	} else {
		ans = strings.Repeat("a", n-1) + "b"
	}
	return ans
}

func main() {
	gas := 3
	fmt.Println(generateTheString(gas))
	gas = 2
	fmt.Println(generateTheString(gas))
	gas = 7
	fmt.Println(generateTheString(gas))
}
