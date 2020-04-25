// 2020/4/18
package main

import (
	"fmt"
	// "math"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
	nums = []int{1, 0, 1, 1}
	k = 1
	fmt.Println(containsNearbyDuplicate(nums, k))
	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Println(!containsNearbyDuplicate(nums, k))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	dic := make(map[int]int)
	for i, v := range nums {
		if _, ok := dic[v]; ok && i-dic[v] <= k {
			return true
		} else {
			dic[v] = i
		}
	}
	return false
}
