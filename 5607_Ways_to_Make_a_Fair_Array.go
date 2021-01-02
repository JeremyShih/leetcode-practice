// 2020/11/22
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	nums := []int{2, 1, 6, 4}
	//fmt.Println(waysToMakeFair(nums))
	fmt.Println(waysToMakeFair(nums) == 1)
	nums = []int{1, 1, 1}
	fmt.Println(waysToMakeFair(nums) == 3)
	nums = []int{1, 2, 3}
	fmt.Println(waysToMakeFair(nums) == 0)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func waysToMakeFair(nums []int) int {
	ans := 0
	odd, even := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even += nums[i]
		} else {
			odd += nums[i]
		}
	}
	leftOdd, leftEven, cur := 0, 0, 0

	for r := 0; r < len(nums); r++ {
		//ints := []int{}
		cur = nums[r]
		isEven := r%2 == 0
		//fmt.Printf("odd: %d, even: %d, leftOdd: %d, leftEven: %d\n", odd, even, leftOdd, leftEven)
		//fmt.Println(ints)
		//odd, even := 0, 0
		//for i := 0; i < len(ints); i++ {
		//	if i%2 == 0 {
		//		even += ints[i]
		//	} else {
		//		odd += ints[i]
		//	}
		//}
		if isEven {
			even -= cur
		} else {
			odd -= cur
		}

		if even+leftOdd == odd+leftEven {
			ans += 1
		}

		if isEven {
			leftEven += cur
		} else {
			leftOdd += cur
		}
	}
	return ans
}
