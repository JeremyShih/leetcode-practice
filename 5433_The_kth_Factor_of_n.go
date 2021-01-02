// 2020/6/27
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(kthFactor(12, 3) == 3)
	fmt.Println(kthFactor(7, 2) == 7)
	fmt.Println(kthFactor(4, 4) == -1)
	fmt.Println(kthFactor(1, 1) == 1)
	fmt.Println(kthFactor(1000, 3) == 4)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func kthFactor(n int, k int) int {
	count := 1
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if count == k {
				return i
			}
			count++
		}
	}
	return -1
}
