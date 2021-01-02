// 2020/5/31
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{3, 4, 5, 2}
	fmt.Println(maxProduct(nums) == 12)
	nums = []int{1, 5, 4, 5}
	fmt.Println(maxProduct(nums) == 16)
	nums = []int{3, 7}
	fmt.Println(maxProduct(nums) == 12)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxProduct(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i]-1)*(nums[j]-1) > ans {
				ans = (nums[i] - 1) * (nums[j] - 1)
			}
		}
	}
	return ans
}
