// 2020/4/12
package main

import (
	"fmt"
	// "strings"
	"time"
)

func main() {
	start := time.Now()
	s := []int{3, 1, 2, 1}
	m := 5
	fmt.Println(Equal(processQueries(s, m), []int{2, 1, 2, 1}))
	s = []int{4, 1, 2, 2}
	m = 4
	fmt.Println(Equal(processQueries(s, m), []int{3, 1, 2, 0}))
	s = []int{7, 5, 5, 8, 3}
	m = 8
	fmt.Println(Equal(processQueries(s, m), []int{6, 5, 0, 7, 5}))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func processQueries(queries []int, m int) []int {
	var ans []int
	p := []int{}
	for i := 1; i <= m; i++ {
		p = append(p, i)
	}
	for _, v := range queries {
		index := SliceIndex(len(p), func(i int) bool { return p[i] == v })
		ans = append(ans, index)
		tmp := p[index]
		copy(p[1:index+1], p[:index])
		p[0] = tmp
	}
	return ans
}

// SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 })
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
