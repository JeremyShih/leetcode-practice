// 2020/7/19
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums) == 167)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxCoins(nums []int) int {
	n, dp := []int{1}, [][]int{}
	n = append(n, nums...)
	n = append(n, 1)
	l := len(n)
	for i := 0; i < l; i++ {
		tmp := make([]int, l)
		dp = append(dp, tmp)
	}

	for gap := 2; gap < l; gap++ {
		for left := 0; left < l-gap; left++ {
			right := left + gap
			for i := left + 1; i < right; i++ {
				if dp[left][right] < dp[left][i]+dp[i][right]+n[left]*n[i]*n[right] {
					dp[left][right] = dp[left][i] + dp[i][right] + n[left]*n[i]*n[right]
				}
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[0][l-1]
}
