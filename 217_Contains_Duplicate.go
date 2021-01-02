// 2020/4/21
package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main() {
	start := time.Now()
	s1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(s1))
	s1 = []int{1, 2, 3, 4}
	fmt.Println(!containsDuplicate(s1))
	s1 = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(s1))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, n := range nums {
		if v, _ := m[n]; v {
			return true
		} else {
			m[n] = true
		}
	}
	return false
}
