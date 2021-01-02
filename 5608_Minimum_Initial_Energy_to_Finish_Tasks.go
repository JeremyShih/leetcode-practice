// 2020/11/22
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tasks := [][]int{{1, 2}, {2, 4}, {4, 8}}
	fmt.Println(minimumEffort(tasks) == 8)
	tasks = [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}
	fmt.Println(minimumEffort(tasks) == 32)
	tasks = [][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}}
	fmt.Println(minimumEffort(tasks) == 27)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minimumEffort(tasks [][]int) int {
	return 0
}
