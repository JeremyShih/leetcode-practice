// 2020/8/9
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s := "leEeetcode"
	fmt.Println(makeGood(s) == "leetcode")
	s = "abBAcC"
	fmt.Println(makeGood(s) == "")
	s = "s"
	fmt.Println(makeGood(s) == "s")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func makeGood(s string) string {
	if len(s) == 0 {
		return s
	}
	//r := []rune(s)
	i := 0
	for i < len(s)-1 {
		//fmt.Println(string(r[i]), string(r[i+1]))
		a, b := string(s[i]), string(s[i+1])
		if a != b && (strings.ToLower(a) == b || a == strings.ToLower(b)) {
			s = s[0:i] + s[i+2:]
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return s
}
