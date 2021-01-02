// 2020/6/4
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	i := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(i) == 4)
	i = []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(i) == 6)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	var dp []int
	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	//fmt.Println(dp)

	maxLen := 0
	for _, v := range dp {
		if v > maxLen {
			maxLen = v
		}
	}
	return maxLen
}
