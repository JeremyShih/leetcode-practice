// 2020/5/7
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := "babad"
	fmt.Println(s, longestPalindrome(s) == "bab")
	s = "cbbd"
	fmt.Println(s, longestPalindrome(s) == "bb")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	dp, ans, maxLen, n := [][]int{}, string(s[0]), 1, len(s)
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			if i == j {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, 0)
			}
		}
		dp = append(dp, tmp)
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
			maxLen = 2
			ans = s[i : i+2]
		}
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i < j-1; i++ {
			if s[i] == s[j] && dp[i+1][j-1] == 1 {
				//which meas substring s[i:j+1] has the same character at first and last index
				//and the substring in the middle (s[i+1:j]) is also a palindromic substring
				fmt.Println(i, j, i+1, j-1)
				dp[i][j] = 1
				if maxLen < j-i+1 {
					ans = s[i : j+1]
					maxLen = j - i + 1
				}
			}
		}
	}

	for _, v := range dp {
		fmt.Println(v)
	}
	return ans
}
