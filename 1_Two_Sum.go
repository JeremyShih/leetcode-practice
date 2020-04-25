// 2020/4/19
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	nums, target, ans := []int{2, 7, 11, 15}, 9, []int{0, 1}
	fmt.Println(reflect.DeepEqual(twoSum(nums, target), ans))
	nums, target, ans = []int{2, 7, 11, 15}, 18, []int{1, 2}
	fmt.Println(reflect.DeepEqual(twoSum(nums, target), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func twoSum(nums []int, target int) []int {
	dic := map[int]int{}
	for i, n := range nums {
		if v, ok := dic[target-n]; ok {
			return []int{v, i}
		}
		dic[n] = i
	}
	return []int{}
}
