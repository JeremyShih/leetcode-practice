// 2020/5/31
package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	h, w, horizontalCuts, verticalCuts := 5, 4, []int{1, 2, 4}, []int{1, 3}
	fmt.Println(maxArea(h, w, horizontalCuts, verticalCuts) == 4)
	h, w, horizontalCuts, verticalCuts = 5, 4, []int{3, 1}, []int{1}
	fmt.Println(maxArea(h, w, horizontalCuts, verticalCuts) == 6)
	h, w, horizontalCuts, verticalCuts = 5, 4, []int{3}, []int{3}
	fmt.Println(maxArea(h, w, horizontalCuts, verticalCuts) == 9)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	//horizontalCuts = append(horizontalCuts, []int{0, h}...)
	//verticalCuts = append(verticalCuts, []int{0, w}...)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	//fmt.Println(horizontalCuts)
	//fmt.Println(verticalCuts)
	ans, maxH, maxV := 0, 0, 0
	for i := 0; i < len(horizontalCuts)-1; i++ {
		diff := math.Mod(float64(horizontalCuts[i+1]-horizontalCuts[i]), math.Pow10(9)+7)
		if int(diff) > maxH {
			maxH = int(diff)
		}
	}
	//fmt.Println(maxH)
	m, n := getMaxMin(horizontalCuts)
	if h-m > maxH {
		maxH = h - m
	}
	if n > maxH {
		maxH = n
	}
	//fmt.Println(maxH)
	for i := 0; i < len(verticalCuts)-1; i++ {
		diff := math.Mod(float64(verticalCuts[i+1]-verticalCuts[i]), math.Pow10(9)+7)
		if int(diff) > maxV {
			maxV = int(diff)
		}
	}
	//fmt.Println(maxV)
	m, n = getMaxMin(verticalCuts)
	//fmt.Println(m, n, w, maxV)
	if w-m > maxV {
		maxV = w - m
	}
	if n > maxV {
		maxV = n
	}
	//fmt.Println(maxH, maxV)
	ans = maxV * maxH
	//for i := 0; i < len(horizontalCuts)-1; i++ {
	//	for j := 0; j < len(verticalCuts)-1; j++ {
	//		area := math.Mod(float64((horizontalCuts[i+1]-horizontalCuts[i])*(verticalCuts[j+1]-verticalCuts[j])), (math.Pow10(9) + 7))
	//		if int(area) > ans {
	//			ans = int(area)
	//		}
	//	}
	//}
	return ans
}

func getMaxMin(a []int) (int, int) {
	m, n := math.MinInt32, math.MaxInt32
	for _, v := range a {
		if v > m {
			m = v
		}
		if v < n {
			n = v
		}
	}
	return m, n
}
