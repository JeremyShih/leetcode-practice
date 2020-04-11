// 2020/3/22
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	fmt.Println(equal(createTargetArray(nums, index), []int{0, 4, 1, 3, 2}))
	nums = []int{1, 2, 3, 4, 0}
	index = []int{0, 1, 2, 3, 0}
	fmt.Println(equal(createTargetArray(nums, index), []int{0, 1, 2, 3, 4}))
	nums = []int{1}
	index = []int{0}
	fmt.Println(equal(createTargetArray(nums, index), []int{1}))

}

func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		res = append(res, 0)
		copy(res[index[i]+1:], res[index[i]:])
		res[index[i]] = nums[i]
	}
	return res
}

func equal(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
