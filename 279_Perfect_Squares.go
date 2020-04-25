// 2020/4/12
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	a, b := 3, 4
	c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
	fmt.Println(a, b, c)

	n := 12
	// fmt.Println(numSquares(n))
	fmt.Println(numSquares(n) == 3)
	n = 13
	fmt.Println(numSquares(n) == 2)
	n = 1
	// fmt.Println(numSquares(n))
	fmt.Println(numSquares(n) == 1)
	n = 2
	fmt.Println(numSquares(n) == 2)
	n = 3
	fmt.Println(numSquares(n) == 3)
	n = 4
	fmt.Println(numSquares(n) == 1)
	n = 5
	fmt.Println(numSquares(n) == 2)
	n = 6
	fmt.Println(numSquares(n) == 3)
	n = 7
	fmt.Println(numSquares(n) == 4)
	for i := 1; i < 589; i++ {
		n += numSquares(i)
	}
	fmt.Println(n)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numSquares(n int) int {
	dp := []int{}
	for i := 0; i <= n; i++ {
		dp = append(dp, math.MaxInt32)
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		m, j := math.MaxInt32, 1
		for i-j*j >= 0 {
			if dp[i-j*j]+1 < m {
				m = dp[i-j*j] + 1
			}
			j++
		}
		dp[i] = m

	}
	return dp[n]
}

// func numSquares(n int) int {
// 	sqr := math.Sqrt(float64(n))
// 	// fmt.Println(sqr)
// 	var nums, test []int
// 	for i := 1; i*i <= n; i++ {
// 		if n == i*i {
// 			return 1
// 		}
// 		nums = append(nums, i*i)
// 	}
// 	for i := 0; i < int(sqr*0.71)+1; i++ {
// 		test = append(test, i*i)
// 	}
// 	// fmt.Println(nums, test)
// 	for _, i := range test {
// 		for _, j := range test {
// 			for _, k := range nums {
// 				if n-i-j == k {
// 					l, m := 0, 0
// 					if i == 0 {
// 						l = 1
// 					}
// 					if j == 0 {
// 						m = 1
// 					}
// 					return 3 - l - m
// 				}
// 			}
// 		}
// 	}
// 	return 4
// }
