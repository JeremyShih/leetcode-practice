// 2020/7/5
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	n, left, right := 4, []int{4, 3}, []int{0, 1}
	fmt.Println(getLastMoment(n, left, right) == 4)
	n, left, right = 7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(getLastMoment(n, left, right) == 7)
	n, right, left = 7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(getLastMoment(n, left, right) == 7)
	n, left, right = 9, []int{5}, []int{4}
	fmt.Println(getLastMoment(n, left, right) == 5)
	n, left, right = 6, []int{6}, []int{0}
	fmt.Println(getLastMoment(n, left, right) == 6)
	n, left, right = 20, []int{4, 7, 15}, []int{9, 3, 13, 10}
	fmt.Println(getLastMoment(n, left, right) == 17)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func getLastMoment(n int, left []int, right []int) int {
	if len(left) == 0 && len(right) == 0 {
		return 0
	}
	if len(left) == 0 {
		sort.Ints(right)
		return n - right[0]
	}
	if len(right) == 0 {
		sort.Ints(left)
		return left[len(left)-1]
	}
	sort.Ints(left)
	sort.Ints(right)
	ans := left[len(left)-1]
	if n-right[0] > ans {
		ans = n - right[0]
	}
	return ans
}
