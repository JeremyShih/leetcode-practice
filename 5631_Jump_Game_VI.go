// 2020/12/20
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums, k := []int{1, -1, -2, 4, -7, 3}, 2
	fmt.Println(maxResult(nums, k) == 7)
	//nums, k = []int{10, -5, -2, 4, 0, 3}, 3
	//fmt.Println(maxResult(nums, k) == 17)
	//nums, k = []int{1, -5, -20, 4, -1, 3, -6, -3}, 2
	//fmt.Println(maxResult(nums, k) == 0)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxResult(nums []int, k int) int {
	//cur := 0
	dp := make([]int, len(nums))
	m := make(map[int]int)
	fmt.Println(nums, k)
	dp[0] = nums[0]
	prevMax := dp[0]
	m[dp[0]] = 0
	for i := 1; i < len(dp); i++ {
		start := i - k
		if start < 0 {
			start = 0
		}

		dp[i] = prevMax
		dp[i] += nums[i]
		m[dp[i]] = i

		if dp[i] > prevMax {
			prevMax = dp[i]
		}

		if m[prevMax] < start {
			delete(m, prevMax)
			prevMax = dp[start]
			for j := start + 1; j <= i; j++ {
				if dp[j] > prevMax {
					prevMax = dp[j]
				}
			}
		}

		fmt.Println(dp, prevMax)
	}
	//fmt.Println(dp)
	return dp[len(dp)-1]
}
