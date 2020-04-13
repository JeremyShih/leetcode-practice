// 2020/4/12
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n := 12
	fmt.Println(numSquares(n) == 3)
	n = 13
	fmt.Println(numSquares(n) == 2)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numSquares(n int) int {
	nums := []int{}
	for i := 1; i*i < n; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
	return 0
}
