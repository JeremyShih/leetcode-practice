// 2020/7/26
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	target := "10111"
	fmt.Println(minFlips(target) == 3)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minFlips(target string) int {
	ans := 0
	s := []rune(target)
	if s[0]-'0' == 1 {
		ans += 1
	}
	for i := 1; i < len(target); i++ {
		if s[i] != s[i-1] {
			ans += 1
		}
	}
	return ans
}
