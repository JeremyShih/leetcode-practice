// 2020/10/18
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, a, b := "5525", 9, 2
	//fmt.Println(strings.EqualFold(findLexSmallestString(s, a, b), "2050"))
	//s, a, b = "74", 5, 1
	//fmt.Println(strings.EqualFold(findLexSmallestString(s, a, b), "24"))
	//s, a, b = "0011", 4, 2
	//fmt.Println(strings.EqualFold(findLexSmallestString(s, a, b), "0011"))
	s, a, b = "43987654", 7, 3
	fmt.Println(strings.EqualFold(findLexSmallestString(s, a, b), "00553311"))
	//fmt.Println(operateA(s, a))
	fmt.Println(operateB(s, b))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func findLexSmallestString(s string, a int, b int) string {
	before, after := "temp", s
	for !strings.EqualFold(before, after) {
		before = after
		after = ""
		resultA, resultB := operateA(before, a), operateB(before, b)
		if !strings.EqualFold(resultA, resultB) {
			if smallerThan(resultA, resultB) {
				after = resultA
			} else {
				after = resultB
			}
		} else {
			if smallerThan(resultA, before) {
				after = resultA
			} else {
				after = before
			}
		}

		if !smallerThan(after, before) {
			after = before
		}
	}
	fmt.Println(before)
	return before
}

func operateA(s string, a int) string {
	r := []int32{}
	for _, c := range s {
		r = append(r, c-'0')
	}
	for i := 0; i < len(r); i++ {
		c := r[i]
		if i%2 == 1 {
			c = (c + int32(a)) % 10
		}
		r[i] = c
	}

	t := []string{}
	for _, c := range r {
		t = append(t, strconv.Itoa(int(c)))
	}
	return strings.Join(t, "")
}

func operateB(s string, b int) string {
	l := len(s)
	t := []string{}
	for i := b; i < l; i += b {
		t = append(t, s[l-i:]+s[:l-i])
	}
	smallest := t[0]
	for i := 1; i < len(t); i++ {
		if smallerThan(t[i], smallest) {
			smallest = t[i]
		}
	}
	return smallest
}

func smallerThan(a, b string) bool {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	return x < y
}
