// 2020/5/17
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	text := "Leetcode is cool"
	ans := "Is cool leetcode"
	fmt.Println(arrangeWords(text) == ans)
	text = "Keep calm and code on"
	ans = "On and keep calm code"
	fmt.Println(arrangeWords(text) == ans)
	text = "To be or not to be"
	ans = "To be or to be not"
	fmt.Println(arrangeWords(text) == ans)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func arrangeWords(text string) string {
	s := strings.Split(strings.ToLower(text), " ")
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if len(s[j-1]) > len(s[j]) {
				tmp := s[j]
				s[j] = s[j-1]
				s[j-1] = tmp
			}
		}
	}
	text = strings.Join(s, " ")
	//f := text[0]
	text = strings.ToUpper(string(text[0])) + text[1:]
	//fmt.Println(strings.ToUpper(string(f)))
	return text
}
