// 2020/3/29
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 5, 3, 4, 1}
	fmt.Println(numTeams(nums) == 3)
	nums = []int{2, 1, 3}
	fmt.Println(numTeams(nums) == 0)
}

func numTeams(rating []int) int {
	if len(rating) < 3 {
		return 0
	}
	count := 0
	for i := 0; i < len(rating)-2; i++ {
		for j := i + 1; j < len(rating)-1; j++ {
			for k := j + 1; k < len(rating); k++ {
				if (rating[i] > rating[j] && rating[j] > rating[k]) || (rating[i] < rating[j] && rating[j] < rating[k]) {
					count += 1
				}
			}
		}
	}
	return count
}
