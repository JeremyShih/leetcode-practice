// 2020/4/26
package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main() {
	start := time.Now()
	s := "011101"
	fmt.Println(maxScore(s) == 5)
	s = "00111"
	fmt.Println(maxScore(s) == 5)
	s = "00"
	// fmt.Println(maxScore(s))
	fmt.Println(maxScore(s) == 1)
	s = "1111"
	fmt.Println(maxScore(s) == 3)
	s = "1"
	for i := 0; i < 498; i++ {
		s += "1"
	}
	// fmt.Println(maxScore(s))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxScore(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		l := s[:i]
		r := s[i:]
		score := 0
		// fmt.Println(l, r)
		for _, c := range l {
			if c-'0' == 0 {
				score++
			}
		}
		for _, c := range r {
			if c-'1' == 0 {
				score++
			}
		}
		if score > ans {
			ans = score
		}
	}
	return ans
}
