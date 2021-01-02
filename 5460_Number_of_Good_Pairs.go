// 2020/7/12
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

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}
