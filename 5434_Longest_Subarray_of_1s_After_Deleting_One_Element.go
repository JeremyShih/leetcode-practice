// 2020/6/27
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{1, 1, 0, 1}
	fmt.Println(longestSubarray(nums) == 3)
	nums = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums) == 5)
	nums = []int{1, 1, 1}
	fmt.Println(longestSubarray(nums) == 2)
	nums = []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums) == 4)
	nums = []int{0, 0, 0}
	fmt.Println(longestSubarray(nums) == 0)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func longestSubarray(nums []int) int {
	p, tmp := []int{}, 0
	for _, v := range nums {
		if v == 0 {
			if tmp != 0 {
				p = append(p, tmp)
				tmp = 0
			}
			p = append(p, v)
		} else {
			tmp += v
		}
	}
	if tmp != 0 {
		p = append(p, tmp)
	}
	//fmt.Println(p)
	if len(p) == 1 {
		return p[0] - 1
	}

	maxLen := 0
	for i, v := range p {
		if v == 0 && i > 0 && i < len(p)-1 {
			if p[i-1]+p[i+1] > maxLen {
				maxLen = p[i-1] + p[i+1]
			}
		}
	}
	return maxLen
}
