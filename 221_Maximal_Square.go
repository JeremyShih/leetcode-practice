// 2020/6/6
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(input) == 4)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maximalSquare(matrix [][]byte) int {
	var dp [][]int
	h := len(matrix)
	if h == 0 {
		return 0
	}

	w := len(matrix[0])
	if w == 0 {
		return 0
	}
	for i := 0; i < h; i++ {
		var tmp []int
		for j := 0; j < w; j++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}

	res := 0
	if matrix[0][0] == 1 {
		dp[0][0] = 1
	}
	//for _, v := range matrix {
	//	fmt.Println(v)
	//}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j]-'0' == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1]
					if dp[i][j] > dp[i-1][j] {
						dp[i][j] = dp[i-1][j]
					}
					if dp[i][j] > dp[i][j-1] {
						dp[i][j] = dp[i][j-1]
					}
					dp[i][j] += 1
				}
				//fmt.Println(dp[i][j])
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	//for _, v := range dp {
	//	fmt.Println(v)
	//}
	return res * res
}
