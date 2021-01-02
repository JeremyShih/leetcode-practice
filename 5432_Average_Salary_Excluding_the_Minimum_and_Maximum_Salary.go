// 2020/6/27
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(average([]int{300, 400, 500}) == 400)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func average(salary []int) float64 {
	minSal, maxSal, sumSal := salary[0], salary[0], 0
	for _, v := range salary {
		if v < minSal {
			minSal = v
		}
		if v > maxSal {
			maxSal = v
		}
		sumSal += v
	}
	ans := float64(sumSal-minSal-maxSal) / float64(len(salary)-2)
	return ans
}
