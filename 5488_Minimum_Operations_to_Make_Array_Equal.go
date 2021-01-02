// 2020/8/16
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i < 10; i++ {
		fmt.Println(minOperations(i))
	}
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minOperations(n int) int {
	if n%2 == 1 {
		return (n*n - 1) / 4
	} else {
		return n * n / 4
	}
}
