// 2020/8/2
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	arr := []int{0}
	for i := 1; i < 100000; i++ {
		arr = append(arr, i)
	}
	k := 99999
	fmt.Println(getWinner(arr, k))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

var arrLeng int

func getWinner(arr []int, k int) int {
	win := arr[0]
	arrLeng = len(arr)
	if k > arrLeng-1 {
		for i := 1; i < len(arr); i++ {
			if arr[i] > win {
				win = arr[i]
			}
		}
		return win
	}

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			if isMax(arr, i, 0, k) {
				return arr[i]
			}
		} else {
			if isMax(arr, i, 1, k-1) {
				return arr[i]
			}
		}
	}
	return win
}

func isMax(arr []int, index int, left int, right int) bool {
	cur := arr[index]
	for i := 1; i <= left; i++ {
		if arr[index-i] > cur {
			return false
		}
	}
	for i := 1; i <= right; i++ {
		if arr[(index+i)%arrLeng] > cur {
			return false
		}
	}
	return true
}
