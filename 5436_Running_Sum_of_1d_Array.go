// 2020/6/14
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func runningSum(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}
	return nums
}
