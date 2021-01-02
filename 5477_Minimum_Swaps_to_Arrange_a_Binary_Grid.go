// 2020/8/2
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	grid := [][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}
	fmt.Println(minSwaps(grid) == 3)
	grid = [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}
	fmt.Println(minSwaps(grid) == -1)
	grid = [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}
	fmt.Println(minSwaps(grid) == 0)
	grid = [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}
	fmt.Println(minSwaps(grid) == 2)
	grid = [][]int{{0, 1, 1, 0}, {1, 1, 1, 0}, {1, 1, 1, 0}, {1, 0, 0, 0}}
	fmt.Println(minSwaps(grid) == -1)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

var n int

func minSwaps(grid [][]int) int {
	ans := 0
	n = len(grid)
	w := len(grid)
	if isValid(grid) {
		return 0
	}

	for i := 0; i < w-1; i++ {
		row := 0
		for row < n && !isValidRow(grid[row], i+1) {
			row++
		}
		//fmt.Println(row)
		if row == n {
			return -1
		}
		ans += row
		n -= 1
		if row+1 <= len(grid) {
			copy(grid[row:], grid[row+1:])
		}
		grid = grid[:len(grid)-1]
		//fmt.Println(grid)
	}
	return ans
}

func isValidRow(row []int, index int) bool {
	for i := index; i < len(row); i++ {
		if row[i] != 0 {
			return false
		}
	}
	return true
}

func isValid(grid [][]int) bool {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
