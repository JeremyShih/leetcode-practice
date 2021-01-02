// 2020/5/24
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	sentence := "i love eating burger"
	searchWord := "burg"
	fmt.Println(isPrefixOfWord(sentence, searchWord) == 4)
	sentence = "this problem is an easy problem"
	searchWord = "pro"
	fmt.Println(isPrefixOfWord(sentence, searchWord) == 2)
	sentence = "i am tired"
	searchWord = "you"
	fmt.Println(isPrefixOfWord(sentence, searchWord) == -1)
	sentence = "i use triple pillow"
	searchWord = "pill"
	fmt.Println(isPrefixOfWord(sentence, searchWord) == 4)
	sentence = "hello from the other side"
	searchWord = "they"
	fmt.Println(isPrefixOfWord(sentence, searchWord) == -1)
	sentence = "hellohello hellohellohello"
	searchWord = "ell"
	fmt.Println(isPrefixOfWord(sentence, searchWord) == -1)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func isPrefixOfWord(sentence string, searchWord string) int {
	s, ans := strings.Split(sentence, " "), -1
	for i, t := range s {
		if a := strings.Index(t, searchWord); a == 0 {
			return i + 1
		}
	}

	//fmt.Println(s)
	return ans
}
