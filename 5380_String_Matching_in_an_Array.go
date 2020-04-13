// 2020/4/12
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(Equal(stringMatching(s), []string{"as", "hero"}))
	s = []string{"leetcode", "et", "code"}
	// fmt.Println(stringMatching(s))
	fmt.Println(Equal(stringMatching(s), []string{"et", "code"}))
	s = []string{"blue", "green", "bu"}
	fmt.Println(Equal(stringMatching(s), []string{}))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func Equal(a, b []string) bool {
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

func stringMatching(words []string) []string {
	ans := []string{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			// fmt.Println(words[i], words[j])
			if strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}
