// 2020/4/12
package main

import (
	"fmt"
	// "strings"
	"time"
)

func main() {
	start := time.Now()
	n := 1
	fmt.Println(numOfWays(n) == 12)
	n = 2
	fmt.Println(numOfWays(n) == 54)
	n = 3
	fmt.Println(numOfWays(n) == 246)
	n = 7
	fmt.Println(numOfWays(n) == 106494)
	n = 5000
	fmt.Println(numOfWays(n) == 30228214)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numOfWays(n int) int {
	color2 := 6
	color3 := 6
	tmp := 0
	for i := 2; i <= n; i++ {
		tmp = color3
		color3 = (2*color3 + 2*color2) % 1000000007
		color2 = (2*tmp + 3*color2) % 1000000007
	}
	num := (color3 + color2) % 1000000007
	return num
}
