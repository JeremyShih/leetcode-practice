// 2020/12/20
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{4, 2, 4, 5, 6}
	fmt.Println(maximumUniqueSubarray(nums) == 17)
	nums = []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
	fmt.Println(maximumUniqueSubarray(nums) == 8)
	nums = []int{10000}
	fmt.Println(maximumUniqueSubarray(nums) == 10000)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maximumUniqueSubarray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int)
	l, r, ans := 0, 1, 0
	m[nums[l]] = l

	for r < len(nums) {
		//fmt.Println(m)
		if v, ok := m[nums[r]]; ok {
			//fmt.Println(ok, nums[r], v)
			for i := l; i <= v; i++ {
				delete(m, nums[i])
			}
			l = v + 1
		}
		m[nums[r]] = r
		//fmt.Println(m)
		sum := 0
		for i := l; i <= r; i++ {
			sum += nums[i]
		}
		if sum > ans {
			ans = sum
		}
		//fmt.Println(nums[l:r+1], ans)
		r++
	}
	return ans
}
