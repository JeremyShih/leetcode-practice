// 2020/5/3
package main

import (
	"fmt"
	// "math"
	"time"
)

func main() {
	start := time.Now()
	n := []int{3, 2, 3}
	fmt.Println(majorityElement(n) == 3)
	n = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(n) == 2)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func majorityElement(nums []int) int {
	pass, m := len(nums)/2, map[int]int{}
	for _, n := range nums {
		m[n] += 1
		if m[n] > pass {
			return n
		}
	}
	return 0
}
