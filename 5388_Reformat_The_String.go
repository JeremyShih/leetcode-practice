// 2020/4/19
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	s := "a0b1c2"
	ans := "0a1b2c"
	fmt.Println(reformat(s) == ans)
	s = "leetcode"
	ans = ""
	fmt.Println(reformat(s) == ans)
	s = "1229857369"
	ans = ""
	fmt.Println(reformat(s) == ans)
	s = "covid2019"
	ans = "c2o0v1i9d"
	// fmt.Println(reformat(s))
	fmt.Println(reformat(s) == ans)
	s = "ab123"
	ans = "1a2b3"
	// fmt.Println(reformat(s))
	fmt.Println(reformat(s) == ans)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func reformat(s string) string {
	var alpha, nums []rune
	for _, c := range s {
		if c-'0' < 10 && c-'0' > -1 {
			nums = append(nums, c)
		} else {
			alpha = append(alpha, c)
		}
	}
	ans := ""
	if math.Abs(float64(len(alpha)-len(nums))) > 1 {
		return ""
	} else if len(alpha) <= len(nums) {
		for len(alpha) > 0 {
			ans += string(nums[0])
			ans += string(alpha[0])
			alpha, nums = alpha[1:], nums[1:]
		}
		if len(alpha) < len(nums) {
			ans += string(nums[0])
		}
		return ans
	} else if len(alpha) > len(nums) {
		for len(nums) > 0 {
			ans += string(alpha[0])
			ans += string(nums[0])
			alpha, nums = alpha[1:], nums[1:]
		}
		ans += string(alpha[0])
		return ans
	}
	return s
}
