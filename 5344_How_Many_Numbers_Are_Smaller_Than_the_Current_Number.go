// 2020/3/1
package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	ori := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ori[i] = nums[i]
	}
	sort.Ints(nums)
	ord := make(map[int]int)
	prev := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ord[nums[i]] = i
		} else {
			if nums[i] == nums[i-1] {
				ord[nums[i]] = prev
			} else {
				ord[nums[i]] = i
				prev = i
			}
		}
	}
	res := []int{}
	// fmt.Println(ori)
	// fmt.Println(ord)
	for i := 0; i < len(nums); i++ {
		res = append(res, ord[ori[i]])
	}
	return res
}

func main() {
	inp := []int{8, 1, 2, 2, 3}
	fmt.Println(inp)
	fmt.Println(smallerNumbersThanCurrent(inp))
	inp = []int{6, 5, 4, 8}
	fmt.Println(inp)
	fmt.Println(smallerNumbersThanCurrent(inp))
	inp = []int{7,7,7,7}
	fmt.Println(inp)
	fmt.Println(smallerNumbersThanCurrent(inp))
}
