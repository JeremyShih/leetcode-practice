// 2020/5/10
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	arr := []int{2, 3, 1, 6, 7}
	fmt.Println(countTriplets(arr) == 4)
	arr = []int{1, 1, 1, 1, 1}
	fmt.Println(countTriplets(arr) == 10)
	arr = []int{2, 3}
	fmt.Println(countTriplets(arr) == 0)
	arr = []int{1, 3, 5, 7, 9}
	fmt.Println(countTriplets(arr) == 3)
	arr = []int{7, 11, 12, 9, 5, 2, 7, 17, 22}
	fmt.Println(countTriplets(arr) == 8)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func countTriplets(arr []int) int {
	triplets := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j; k < len(arr); k++ {
				a, b := arr[i], arr[j]
				for l := i + 1; l < j; l++ {
					a = a ^ arr[l]
				}
				for m := j + 1; m <= k; m++ {
					b = b ^ arr[m]
				}
				if a == b {
					triplets = append(triplets, []int{i, j, k})
				}
			}
		}
	}
	//fmt.Println(triplets)
	return len(triplets)
}
