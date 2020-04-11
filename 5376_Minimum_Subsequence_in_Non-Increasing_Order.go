// 2020/4/5
package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4, 3, 10, 9, 8}
	ans := []int{10, 9}
	fmt.Println(Equal(minSubsequence(nums), ans))
	nums = []int{4, 4, 7, 6, 7}
	ans = []int{7, 7, 6}
	fmt.Println(Equal(minSubsequence(nums), ans))
	nums = []int{6}
	ans = []int{6}
	fmt.Println(Equal(minSubsequence(nums), ans))
}

func minSubsequence(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	res := []int{}
	for Sum(nums) >= Sum(res) {
		i, m := Max(nums)
		res = append(res, m)
		copy(nums[i:], nums[i+1:])
		nums = nums[:len(nums)-1]
		// fmt.Println(nums)
	}
	// fmt.Println(res)
	return res
}

func Sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func Max(a []int) (int, int) {
	m := math.MinInt32
	index := -1
	for i, v := range a {
		if v > m {
			m = v
			index = i
		}
	}
	return index, m
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
