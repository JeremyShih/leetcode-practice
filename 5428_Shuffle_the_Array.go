// 2020/6/7
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	ans := []int{2, 3, 5, 4, 1, 7}
	fmt.Println(reflect.DeepEqual(shuffle(nums, n), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func shuffle(nums []int, n int) []int {
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i])
		ans = append(ans, nums[i+n])
	}
	return ans
}
