// 2020/4/12
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	t := 5
	fmt.Println(searchMatrix(matrix, t))
	t = 20
	fmt.Println(!searchMatrix(matrix, t))
	matrix = [][]int{{}}
	fmt.Println(matrix, !searchMatrix(matrix, t))
	matrix = [][]int{{}, {}}
	fmt.Println(matrix, !searchMatrix(matrix, t))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m < 1 {
		return false
	}
	n := len(matrix[0])
	if n < 1 {
		return false
	}
	row, col := 0, n-1
	for row < m && col > -1 {
		cur := matrix[row][col]
		// fmt.Println(cur)
		if target > cur {
			row++
		} else if target < cur {
			col--
		} else {
			return true
		}
	}
	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	m := len(matrix)
// 	if m < 1 {
// 		return false
// 	}
// 	n := len(matrix[0])
// 	if n < 1 {
// 		return false
// 	}
// 	i := 0
// 	for i < m && target >= matrix[i][0] {
// 		left, right := 0, n
// 		for left < right {
// 			mid := (left + right) / 2
// 			if target > matrix[i][mid] {
// 				left = mid + 1
// 			} else if target < matrix[i][mid] {
// 				right = mid
// 			} else {
// 				return true
// 			}
// 		}
// 		i++
// 	}
// 	return false
// }
