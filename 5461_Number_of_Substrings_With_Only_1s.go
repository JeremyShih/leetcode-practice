// 2020/7/12
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	s := "0110111"
	fmt.Println(numSub(s) == 9)
	s = "101"
	fmt.Println(numSub(s) == 2)
	s = "101"
	fmt.Println(numSub(s) == 2)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numSub(s string) int {
	ans := 0
	var dp []int
	for i := 0; i < len(s); i++ {
		dp = append(dp, 0)
	}
	for i := 0; i < len(s); i++ {
		if s[i]-'0' == 1 {
			dp[i] = 1
			if i > 0 && dp[i] > 0 {
				dp[i] += dp[i-1]
			}
			ans += dp[i]
			ans = ans % int(math.Pow10(7)+9)
		}
	}
	return ans
}
