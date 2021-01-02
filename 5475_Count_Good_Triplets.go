// 2020/8/2
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3
	//fmt.Println(countGoodTriplets(arr, a, b, c) == 4)
	//arr = []int{1, 1, 2, 2, 3}
	//a, b, c = 0, 0, 1
	//fmt.Println(countGoodTriplets(arr, a, b, c) == 0)
	arr = []int{7, 3, 7, 3, 12, 1, 12, 2, 3}
	a, b, c = 5, 8, 1
	fmt.Println(countGoodTriplets(arr, a, b, c))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ans := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				//fmt.Print([]int{arr[i] - arr[j], arr[j] - arr[k], arr[i] - arr[k]})
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) && math.Abs(float64(arr[j]-arr[k])) <= float64(b) && math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					//fmt.Println([]int{arr[i], arr[j], arr[k]})
					ans++
				}
			}
		}
	}
	return ans
}
