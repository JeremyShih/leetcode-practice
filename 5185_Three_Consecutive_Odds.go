// 2020/8/16
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	for i, v := range arr {
		if i+2 == len(arr) {
			break
		}
		if v%2 == 1 {
			if arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
				return true
			}
		}
	}
	return false
}
