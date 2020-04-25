// 2020/4/18
package main

import (
	"fmt"
	// "math"
	"time"
)

func main() {
	start := time.Now()
	n := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(n)
	ans := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	fmt.Println(String2DSliceEqual(n, ans))
	n = [][]int{{0, 1, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 0}}
	ans = [][]int{{0, 0, 0}, {1, 0, 1}, {1, 0, 1}, {1, 0, 1}}
	gameOfLife(n)
	fmt.Println(String2DSliceEqual(n, ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func gameOfLife(board [][]int) {
	tmp := make([][]int, len(board))
	for i := range board {
		tmp[i] = make([]int, len(board[i]))
		copy(tmp[i], board[i])
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = checkNeighbors(tmp, i, j)
		}
	}
}

func checkNeighbors(board [][]int, i, j int) int {
	count, ori := 0, board[i][j]
	for x := i - 1; x <= i+1; x++ {
		if x > -1 && x < len(board) {
			if j > 0 && board[x][j-1] == 1 {
				count++
			}
			if j < len(board[0])-1 && board[x][j+1] == 1 {
				count++
			}
			if x != i && board[x][j] == 1 {
				count++
			}
		}
	}
	if ori == 1 {
		if count < 2 {
			ori = 0
		} else if count > 3 {
			ori = 0
		}
	} else {
		if count == 3 {
			ori = 1
		}
	}
	return ori
}

func String2DSliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if !StringSliceEqual(v, b[i]) {
			return false
		}
	}
	return true
}

func StringSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
