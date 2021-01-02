// 2020/11/22
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	word1, word2 := []string{"ab", "c"}, []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
	word1, word2 = []string{"a", "cb"}, []string{"a", "bc"}
	fmt.Println(!arrayStringsAreEqual(word1, word2))
	word1, word2 = []string{"abc", "d", "defg"}, []string{"abcddefg"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	if strings.EqualFold(strings.Join(word1, ""), strings.Join(word2, "")) {
		return true
	}
	return false
}
