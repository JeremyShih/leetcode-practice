// 2020/11/8
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	n := 7
	fmt.Println(getMaximumGenerated(n) == 3)
	n = 2
	fmt.Println(getMaximumGenerated(n) == 1)
	n = 3
	fmt.Println(getMaximumGenerated(n) == 2)
	n = 0
	fmt.Println(getMaximumGenerated(n) == 0)
	n = 1
	fmt.Println(getMaximumGenerated(n) == 1)
	n = 20
	fmt.Println(getMaximumGenerated(n))
	n = 21
	fmt.Println(getMaximumGenerated(n))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	ans := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[(i-1)/2] + arr[(i+1)/2]
		}
		if arr[i] > ans {
			ans = arr[i]
		}
	}
	return ans
}
