// 2020/4/26
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ans := []int{1, 4, 2, 7, 5, 3, 8, 6, 9}
	fmt.Println(reflect.DeepEqual(findDiagonalOrder(nums), ans))
	nums = [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}
	ans = []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}
	fmt.Println(reflect.DeepEqual(findDiagonalOrder(nums), ans))
	nums = [][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}
	ans = []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11}
	fmt.Println(reflect.DeepEqual(findDiagonalOrder(nums), ans))
	nums = [][]int{{1, 2, 3, 4, 5, 6}}
	ans = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reflect.DeepEqual(findDiagonalOrder(nums), ans))
	nums = [][]int{{6}, {8}, {6, 1, 6, 16}}
	ans = []int{6, 8, 6, 1, 6, 16}
	fmt.Println(reflect.DeepEqual(findDiagonalOrder(nums), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func findDiagonalOrder(nums [][]int) []int {
	ans := []int{}
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if len(nums[i]) > maxLen {
			maxLen = len(nums[i])
		}
		for x := i; x >= 0; x-- {
			y := i - x
			// fmt.Println(x, len(nums))
			if x >= len(nums) {
				continue
			}
			if y >= len(nums[x]) {
				continue
			}
			ans = append(ans, nums[x][y])
			// fmt.Println(x, y)
		}
	}
	// fmt.Println(maxLen)
	for i := len(nums); i < maxLen+len(nums); i++ {
		for x := i; x >= 0; x-- {
			y := i - x
			// fmt.Println(x, len(nums))
			if x >= len(nums) {
				continue
			}
			if y >= len(nums[x]) {
				continue
			}
			ans = append(ans, nums[x][y])
			// fmt.Println(x, y)
		}
	}
	return ans
}
