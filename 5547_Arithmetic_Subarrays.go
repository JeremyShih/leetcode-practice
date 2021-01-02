// 2020/10/25
package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	fmt.Println(reflect.DeepEqual(checkArithmeticSubarrays(nums, l, r), []bool{true, false, true}))
	nums = []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}
	l = []int{0, 1, 6, 4, 8, 7}
	r = []int{4, 4, 9, 7, 9, 10}
	fmt.Println(reflect.DeepEqual(checkArithmeticSubarrays(nums, l, r), []bool{false, true, false, false, true, true}))
	//fmt.Println(isArithmetic(nums[:3]))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := []bool{}
	for i := 0; i < len(l); i++ {
		//fmt.Println(nums)
		c := []int{}
		for j := l[i]; j < r[i]+1; j++ {
			c = append(c, nums[j])
		}
		ans = append(ans, isArithmetic(c))
	}
	//fmt.Println(ans)
	return ans
}

func isArithmetic(s []int) bool {
	//fmt.Println(s)
	if len(s) <= 2 {
		return true
	}
	sort.Ints(s)

	for i := 2; i < len(s); i++ {
		if s[i-1]-s[i-2] != s[i]-s[i-1] {
			return false
		}
	}
	return true
}
