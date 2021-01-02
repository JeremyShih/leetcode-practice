// 2020/10/18
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := "aa"
	fmt.Println(maxLengthBetweenEqualCharacters(s) == 0)
	s = "abca"
	fmt.Println(maxLengthBetweenEqualCharacters(s) == 2)
	s = "cbzxy"
	fmt.Println(maxLengthBetweenEqualCharacters(s) == -1)
	s = "cabbac"
	fmt.Println(maxLengthBetweenEqualCharacters(s) == 4)
	s = "mgntdygtxrvxjnwksqhxuxtrv"
	fmt.Println(maxLengthBetweenEqualCharacters(s) == 18)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxLengthBetweenEqualCharacters(s string) int {
	maxLength := -1
	if len(s) < 2 {
		return maxLength
	}
	m := make(map[rune]int)
	for i, c := range s {
		if v, ok := m[c]; ok {
			distance := i - v - 1
			if distance > maxLength {
				maxLength = distance
			}
		} else {
			m[c] = i
		}
	}
	return maxLength
}
