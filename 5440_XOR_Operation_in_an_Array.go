// 2020/6/21
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(xorOperation(5, 0) == 8)
	fmt.Println(xorOperation(4, 3) == 8)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func xorOperation(n int, start int) int {
	ans := 0
	for i := start; i < start+2*n; i += 2 {
		ans ^= i
	}
	return ans
}
