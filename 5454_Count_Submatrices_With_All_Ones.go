// 2020/7/5
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	mat := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(numSubmat(mat) == 13)
	mat = [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}
	//fmt.Println(numSubmat(mat) == 24)
	//mat = [][]int{{1, 1, 1, 1, 1, 1}}
	//fmt.Println(numSubmat(mat) == 21)
	//mat = [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
	//fmt.Println(numSubmat(mat) == 5)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numSubmat(mat [][]int) int {
	if len(mat) == 0 {
		return 0
	}
	m := len(mat)
	if len(mat[0]) == 0 {
		return 0
	}
	n := len(mat[0])
	dp := [][]int{}
	for i := 0; i < m; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				dp[i][j] = 1
				tmp := math.MaxInt16
				if i > 0 && mat[i-1][j] == 1 {
					if dp[i-1][j] < tmp {
						tmp = dp[i-1][j]
					}
				}
				if j > 0 && mat[i][j-1] == 1 {
					if dp[i][j-1] < tmp {
						tmp = dp[i][j-1]
					}
				}
				if i > 0 && j > 0 && mat[i-1][j] == 1 && mat[i][j-1] == 1 && mat[i-1][j-1] == 1 {
					if dp[i-1][j-1] < tmp {
						tmp = dp[i-1][j-1]
					}
				}
				dp[i][j] += tmp
				ans += dp[i][j]
			}
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	fmt.Println(ans)
	return ans
}
