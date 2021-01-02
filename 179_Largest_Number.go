// 2020/5/3
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums) == "210")
	nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums) == "9534330")
	nums = []int{10, 0}
	fmt.Println(largestNumber(nums) == "100")
	nums = []int{0, 0, 0}
	fmt.Println(largestNumber(nums) == "0")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func largestNumber(nums []int) string {
	var s []string
	n := len(nums)
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(nums[i]))
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			a, b := s[i], s[j]
			l, _ := strconv.Atoi(a + b)
			r, _ := strconv.Atoi(b + a)
			if l < r {
				s[i], s[j] = b, a
			}
		}
	}
	if n > 0 && s[0] == "0" {
		return "0"
	}
	return strings.Join(s[:], "")
}
