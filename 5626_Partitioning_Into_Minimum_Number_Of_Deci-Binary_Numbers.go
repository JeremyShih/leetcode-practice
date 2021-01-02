// 2020/12/13
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n := "32"
	fmt.Println(minPartitions(n) == 3)
	n = "82734"
	fmt.Println(minPartitions(n) == 8)
	n = "27346209830709182346"
	fmt.Println(minPartitions(n) == 9)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minPartitions(n string) int {
	ans := 0
	for _, c := range n {
		if int(c-'0') > ans {
			ans = int(c - '0')
		}
	}
	//fmt.Println(ans)
	return ans
}
