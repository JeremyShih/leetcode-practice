// 2020/11/22
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	n, k := 3, 27
	//fmt.Println(getSmallestString(n, k))
	fmt.Println(strings.EqualFold(getSmallestString(n, k), "aay"))
	n, k = 5, 73
	fmt.Println(strings.EqualFold(getSmallestString(n, k), "aaszz"))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func getSmallestString(n int, k int) string {
	s := []rune{}
	for i := 0; i < n; i++ {
		s = append(s, 1)
	}
	k -= n
	cur := n - 1
	for k > 0 {
		if k+int(s[cur]) > 26 {
			s[cur] = 26
			k -= 25
			cur--
		} else {
			//fmt.Println(k, rune(k))
			s[cur] += rune(k)
			k = 0
		}
	}
	//fmt.Println(s)

	c := make([]string, n)
	for i, v := range s {
		//fmt.Println(v, string(v+'a'-1), "=")
		c[i] = string(v + 'a' - 1)
	}
	//fmt.Println(c)
	return strings.Join(c, "")
}
