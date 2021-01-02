// 2020/12/13
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n := 7
	fmt.Println(numberOfMatches(n) == 6)
	n = 14
	fmt.Println(numberOfMatches(n) == 13)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numberOfMatches(n int) int {
	ans := 0
	for n > 1 {
		if n%2 == 0 {
			ans += n / 2
			n /= 2
		} else {
			ans += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return ans
}
