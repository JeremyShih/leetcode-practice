// 2020/12/6
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	nums, k := []int{1, 2, 3, 4}, 5
	fmt.Println(maxOperations(nums, k) == 2)
	nums, k = []int{3, 1, 3, 4, 3}, 6
	fmt.Println(maxOperations(nums, k) == 1)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	l, r, ans := 0, len(nums)-1, 0
	for l < r {
		sum := nums[l] + nums[r]
		if sum == k {
			l, r = l+1, r-1
			ans++
		} else if sum > k {
			r--
		} else {
			l++
		}
	}
	fmt.Println(ans)
	return ans
}
