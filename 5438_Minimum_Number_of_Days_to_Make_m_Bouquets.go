// 2020/6/14
//not solved
package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	bloomDay, m, k := []int{1, 10, 3, 10, 2}, 3, 1
	fmt.Println(minDays(bloomDay, m, k) == 3)
	bloomDay, m, k = []int{7, 7, 7, 7, 12, 7, 7}, 2, 3
	fmt.Println(minDays(bloomDay, m, k) == 12)
	//bloomDay, m, k = []int{1000000000, 1000000000}, 1, 1
	//fmt.Println(minDays(bloomDay, m, k) == 1000000000)
	//bloomDay, m, k = []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2
	//fmt.Println(minDays(bloomDay, m, k) == 9)
	//bloomDay, m, k = []int{30, 49, 11, 66, 54, 22, 2, 57, 35}, 3, 3
	//fmt.Println(minDays(bloomDay, m, k) == 57)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}
	set, place, bou := make(map[int]struct{}), make(map[int]int), []int{}
	for i, v := range bloomDay {
		set[v] = struct{}{}
		place[v] = i
		bou = append(bou, 0)
	}
	//fmt.Println("place", place)
	l, dayList := len(bloomDay), []int{}
	for k, _ := range set {
		dayList = append(dayList, k)
	}
	sort.Ints(dayList)
	fmt.Println("dayList", dayList)

	for i := 0; i < l; i++ {
		if i >= k-1 {
			if bouquets(bloomDay[i], i, k, bloomDay) {
				bou[i] = 1
			}
		}
	}
	fmt.Println(bou)

	total := 0
	for _, cur := range dayList {
		for i, v := range bloomDay {
			if cur == v {
				total += bou[i]
			}
		}
		if total >= m {
			return cur
		}
	}
	return -1

	//i := 0
	//for m > 0 {
	//	cur := dayList[i]
	//	for j, _ := range bloomDay {
	//		if j+k > len(bloomDay) {
	//			break
	//		}
	//		if bouquets(cur, j, k, bloomDay) {
	//			m -= 1
	//		}
	//	}
	//	i++
	//	fmt.Println(bloomDay)
	//	if i >= len(dayList) {
	//		return -1
	//	}
	//}
	//return dayList[i-1]
}

func bouquets(cur int, i int, k int, bloomDay []int) bool {
	for j := 0; j < k; j++ {
		if bloomDay[i-j] > cur {
			return false
		}
	}

	//for j := 0; j < k; j++ {
	//	bloomDay[i+j] = math.MaxInt16
	//}
	return true
}
