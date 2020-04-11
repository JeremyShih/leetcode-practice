// 2020/4/4
package main

import "fmt"

func main() {
	str := []string{"11110", "11010", "11000", "00000"}
	grid := [][]byte{}
	for _, s := range str {
		grid = append(grid, []byte(s))
	}
	fmt.Println(numIslands(grid) == 1)
	str = []string{"11000", "11000", "00100", "00011"}
	grid = [][]byte{}
	for _, s := range str {
		grid = append(grid, []byte(s))
	}
	fmt.Println(numIslands(grid) == 3)
	str = []string{"111", "010", "111"}
	grid = [][]byte{}
	for _, s := range str {
		grid = append(grid, []byte(s))
	}
	fmt.Println(numIslands(grid) == 1)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != []byte("0")[0] {
				count += 1
				grid = findAdjacent(grid, i, j)
				// fmt.Println(grid)
			}
		}
	}
	return count
}

func findAdjacent(grid [][]byte, x int, y int) [][]byte {
	grid[x][y] = []byte("0")[0]
	if 0 < x && grid[x-1][y] != []byte("0")[0] {
		grid = findAdjacent(grid, x-1, y)
	}
	if len(grid) > x+1 && grid[x+1][y] != []byte("0")[0] {
		grid = findAdjacent(grid, x+1, y)
	}
	if 0 < y && grid[x][y-1] != []byte("0")[0] {
		grid = findAdjacent(grid, x, y-1)
	}
	if len(grid[0]) > y+1 && grid[x][y+1] != []byte("0")[0] {
		grid = findAdjacent(grid, x, y+1)
	}
	return grid
}
