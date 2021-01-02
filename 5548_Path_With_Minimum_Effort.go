// 2020/10/25
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights) == 2)
	heights = [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights) == 1)
	heights = [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	fmt.Println(minimumEffortPath(heights) == 0)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minimumEffortPath(heights [][]int) int {
	passed := [][]bool{}
	
	return 0
}
