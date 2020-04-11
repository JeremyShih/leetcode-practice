// 2020/4/9
package main

import (
	"fmt"
	// "sort"
)

func main() {
	numbers := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(numbers, k) == 5)
	numbers = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(numbers, k) == 4)
}

func findKthLargest(nums []int, k int) int {
	// fmt.Println(nums)
	QuickSortRecursively(nums, 0, len(nums)-1)
	// fmt.Println(nums)
	return nums[len(nums)-k]
}

// 快速排序法(遞增)，使用遞迴。此為用來遞迴呼叫的函數。
func QuickSortRecursively(array []int, start int, end int) {
	if end > start {
		pivot := array[start]

		l := start + 1
		r := end

		for {
			for r > start && array[r] >= pivot {
				r -= 1
			}

			for l <= r && array[l] <= pivot {
				l += 1
			}

			if l < r {
				array[l], array[r] = array[r], array[l]
			} else {
				if r > start {
					array[start], array[r] = array[r], array[start]
				}

				break
			}
		}

		if r > start {
			QuickSortRecursively(array, start, r-1)
		}

		if r < end {
			QuickSortRecursively(array, r+1, end)
		}
	}
}
