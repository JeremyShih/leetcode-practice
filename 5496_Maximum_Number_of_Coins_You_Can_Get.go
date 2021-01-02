// 2020/8/23
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	piles := []int{2, 4, 1, 2, 7, 8}
	fmt.Println(maxCoins(piles) == 9)
	piles = []int{2, 4, 5}
	fmt.Println(maxCoins(piles) == 4)
	piles = []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
	fmt.Println(maxCoins(piles) == 18)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxCoins(piles []int) int {
	sort.Ints(piles)
	//fmt.Println(piles)
	ans := 0
	for i := len(piles) - 2; i >= len(piles)/3; i -= 2 {
		ans += piles[i]
	}
	return ans
}
