// 2020/12/13
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	stones := []int{5, 3, 1, 4, 2}
	fmt.Println(stoneGameVII(stones) == 6)
	stones = []int{7, 90, 5, 1, 100, 10, 10, 2}
	fmt.Println(stoneGameVII(stones) == 122)
	stones = []int{481, 905, 202, 250, 371, 628, 92, 604, 836, 338, 676, 734}
	fmt.Println(stoneGameVII(stones) == 3459)
	stones = []int{792, 195, 697, 271, 743, 51, 836, 322, 135, 550, 399, 182, 988, 25, 395, 254, 480, 931, 513, 772, 798, 102, 110, 915, 794, 330, 597, 220, 789, 462}
	fmt.Println(stoneGameVII(stones) == 9066)
	stones = []int{121, 903, 609, 929, 646, 419, 823, 722, 223, 170, 8, 704, 102, 803, 639, 548, 364, 306, 440, 65, 933, 123, 21, 376, 215, 798, 280, 232, 942, 513, 463, 567, 34, 642, 958, 823, 37, 605, 784}
	fmt.Println(stoneGameVII(stones))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func stoneGameVII(stones []int) int {
	ans := turn(stones, sum(stones), 0, 0, true)
	//fmt.Println(ans)
	return ans
}

func turn(stones []int, sum, a, b int, left bool) int {
	if len(stones) <= 1 {
		return a - b
	}
	fmt.Println(stones, sum, a, b, left)
	l := len(stones)
	if left {
		removeLeft := turn(stones[1:], sum-stones[0], a+sum-stones[0], b, false)
		removeRight := turn(stones[:l-1], sum-stones[l-1], a+sum-stones[l-1], b, false)
		if removeLeft > removeRight {
			return removeLeft
		} else {
			return removeRight
		}
	} else {
		removeLeft := turn(stones[1:], sum-stones[0], a, b+sum-stones[0], true)
		removeRight := turn(stones[:l-1], sum-stones[l-1], a, b+sum-stones[l-1], true)
		if removeLeft > removeRight {
			return removeRight
		} else {
			return removeLeft
		}
	}
}

func sum(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}
