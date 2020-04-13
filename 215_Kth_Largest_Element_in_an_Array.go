// 2020/4/9
package main

import "fmt"

func main() {
	numbers := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(numbers, k) == 5)
	numbers = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(numbers, k) == 4)
}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(array []int, start int, end int) {
	if start < end {
		pivot := array[start]
		l, r := start+1, end
		for {
			for start < r && array[r] >= pivot {
				r -= 1
			}
			for l <= r && array[l] < pivot {
				l += 1
			}

			if l < r {
				array[l], array[r] = array[r], array[l]
			} else {
				if start < r {
					array[start], array[r] = array[r], array[start]
				}
				break
			}
		}
		if r < end {
			quickSort(array, r+1, end)
		}
		if start < r {
			quickSort(array, start, r-1)
		}
	}
}
