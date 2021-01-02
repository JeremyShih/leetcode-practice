// 2020/12/27
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := "book"
	fmt.Println(halvesAreAlike(s))
	s = "textbook"
	fmt.Println(!halvesAreAlike(s))
	s = "MerryChristmas"
	fmt.Println(!halvesAreAlike(s))
	s = "AbCdEfGh"
	fmt.Println(halvesAreAlike(s))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func halvesAreAlike(s string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	a, b := s[:len(s)/2], s[len(s)/2:]
	//fmt.Println(a, b)
	ac, bc := 0, 0
	for _, c := range a {
		for _, v := range vowels {
			if c-v == 0 {
				ac++
			}
		}
	}
	for _, c := range b {
		for _, v := range vowels {
			if c-v == 0 {
				bc++
			}
		}
	}
	if ac == bc {
		return true
	}
	return false
}
