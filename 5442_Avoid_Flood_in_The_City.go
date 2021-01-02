// 2020/6/21
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	rains := []int{1, 2, 3, 4}
	ans := []int{-1, -1, -1, -1}
	fmt.Println(reflect.DeepEqual(avoidFlood(rains), ans))
	rains = []int{1, 2, 0, 0, 2, 1}
	ans = []int{-1, -1, 1, 2, -1, -1}
	fmt.Println(reflect.DeepEqual(avoidFlood(rains), ans))
	rains = []int{1, 2, 0, 1, 2}
	ans = []int{}
	fmt.Println(reflect.DeepEqual(avoidFlood(rains), ans))
	rains = []int{69, 0, 0, 0, 69}
	ans = []int{-1, 69, 1, 1, -1}
	fmt.Println(reflect.DeepEqual(avoidFlood(rains), ans))
	rains = []int{10, 20, 20}
	ans = []int{}
	fmt.Println(reflect.DeepEqual(avoidFlood(rains), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func avoidFlood(rains []int) []int {
	ans := make([]int, len(rains))
	m := make(map[int]int)

	for i := len(rains) - 1; i > -1; i-- {
		if _, ok := m[rains[i]]; ok && rains[i] > 0 {
			return []int{}
		}
		if rains[i] > 0 {
			ans[i] = -1
			m[rains[i]] = 1
		} else {
			top := 1
			for k, _ := range m {
				top = k
				break
			}
			ans[i] = top
			delete(m, top)
		}
	}
	fmt.Println(ans)
	return ans
}
