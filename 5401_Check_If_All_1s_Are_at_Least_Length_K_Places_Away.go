// 2020/5/3
package main

import (
	"fmt"
	// "math"
	"time"
)

func main() {
	start := time.Now()
	nums, k := []int{1, 0, 0, 0, 1, 0, 0, 1}, 2
	fmt.Println(kLengthApart(nums, k))
	nums, k = []int{1, 0, 0, 1, 0, 1}, 2
	fmt.Println(!kLengthApart(nums, k))
	nums, k = []int{1, 1, 1, 1, 1}, 0
	fmt.Println(kLengthApart(nums, k))
	nums, k = []int{1, 0, 0, 0, 1}, 3
	fmt.Println(kLengthApart(nums, k))
	nums, k = []int{0, 1, 0, 1}, 1
	fmt.Println(kLengthApart(nums, k))
	nums, k = []int{0, 1, 0, 0}, 2
	fmt.Println(kLengthApart(nums, k))
	nums, k = []int{1, 0, 0, 0}, 1
	fmt.Println(kLengthApart(nums, k))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func kLengthApart(nums []int, k int) bool {
	distance, last := len(nums)+1, -1
	for i, v := range nums {
		if v == 1 {
			// fmt.Println(i, last, distance)
			if last != -1 {
				if i-last-1 < distance {
					distance = i - last - 1
				}
			}
			last = i
		}
	}
	// fmt.Println(distance, k)
	if distance >= k {
		return true
	} else {
		return false
	}
}
