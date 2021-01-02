// 2020/7/5
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 3 {
		return true
	}
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			return false
		}
	}
	return true
}
