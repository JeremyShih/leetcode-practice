// 2020/5/3
package main

import (
	"fmt"
	// "math"
	"time"
)

func main() {
	start := time.Now()
	s := "A"
	fmt.Println(titleToNumber(s) == 1)
	s = "AB"
	fmt.Println(titleToNumber(s) == 28)
	s = "ZY"
	fmt.Println(titleToNumber(s) == 701)
	s = "AAA"
	fmt.Println(titleToNumber(s) == 703)
	s = "AAZ"
	fmt.Println(titleToNumber(s) == 728)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func titleToNumber(s string) int {
	r, ans := []rune(s), 0
	for _, v := range r {
		ans = characterToNumber(v) + ans*26
	}
	return ans
}

func characterToNumber(c rune) int {
	return int(c - 'A' + 1)
}
