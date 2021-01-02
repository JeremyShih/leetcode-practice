// 2020/11/8
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := "aab"
	fmt.Println(minDeletions(s) == 0)
	s = "aaabbbcc"
	fmt.Println(minDeletions(s) == 2)
	s = "ceabaacb"
	fmt.Println(minDeletions(s) == 2)
	s = "bbcebab"
	fmt.Println(minDeletions(s) == 2)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minDeletions(s string) int {
	m := make(map[rune]int)
	m1 := make(map[int]bool)
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	ans := 0
	for _, v := range m {
		if _, ok := m1[v]; ok {
			// get deletions
			for i := v - 1; i >= 0; i-- {
				if _, ok := m1[i]; !ok {
					ans += v - i
					if i > 0 {
						m1[i] = true
					}
					break
				}
			}
		} else {
			m1[v] = true
		}
	}
	//fmt.Println(m1)
	return ans
}
