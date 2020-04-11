// 2020/3/29
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 3, 4}
	fmt.Println(findLucky(nums) == 2)
	nums = []int{1, 2, 2, 3, 3, 3}
	fmt.Println(findLucky(nums) == 3)
	nums = []int{2, 2, 2, 3, 3}
	fmt.Println(findLucky(nums) == -1)
	nums = []int{5}
	fmt.Println(findLucky(nums) == -1)
	nums = []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(findLucky(nums) == 7)
}

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	// fmt.Println(m)
	var lucky []int
	for k, v := range m {
		if k == v {
			lucky = append(lucky, v)
		}
	}
	// fmt.Println(lucky)
	max := -1
	for _, n := range lucky {
		if n > max {
			max = n
		}
	}
	return max
}
