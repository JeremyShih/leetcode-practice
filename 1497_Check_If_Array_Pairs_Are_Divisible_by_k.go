// 2020/6/29
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	arr, k := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5
	fmt.Println(canArrange(arr, k))
	arr, k = []int{1, 2, 3, 4, 5, 6}, 7
	fmt.Println(canArrange(arr, k))
	arr, k = []int{1, 2, 3, 4, 5, 6}, 10
	fmt.Println(!canArrange(arr, k))
	arr, k = []int{-10, 10}, 2
	fmt.Println(canArrange(arr, k))
	arr, k = []int{-1, 1, -2, 2, -3, 3, -4, 4}, 3
	fmt.Println(canArrange(arr, k))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func canArrange(arr []int, k int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		r := (v%k + k) % k
		if _, ok := m[r]; ok {
			m[r] += 1
		} else {
			m[r] = 1
		}
	}
	//fmt.Println(m)
	if v, ok := m[0]; ok && v%2 != 0 {
		return false
	}
	delete(m, 0)

	for key, v := range m {
		if val, ok := m[k-key]; key != 0 && !ok {
			return false
		} else if v != val {
			return false
		}
		delete(m, key)
		delete(m, k-key)
	}
	return true
}
