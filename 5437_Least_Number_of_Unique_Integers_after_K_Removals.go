// 2020/6/14
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	arr, k := []int{5, 5, 4}, 1
	fmt.Println(findLeastNumOfUniqueInts(arr, k) == 1)
	arr, k = []int{4, 3, 1, 1, 3, 3, 2}, 3
	fmt.Println(findLeastNumOfUniqueInts(arr, k) == 2)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	//{ key: 1, value: count of int == 1,
	//  key: 2, value: count of int == 2}
	countMap := make(map[int]int)
	for _, v := range m {
		if _, ok := countMap[v]; ok {
			countMap[v] += 1
		} else {
			countMap[v] = 1
		}
	}
	fmt.Println("m:", m)
	fmt.Println("countMap", countMap)

	count := 1
	for k > 0 {
		if _, ok := countMap[count]; ok {
			if countMap[count] > k/count {
				countMap[count] -= k / count
				k = 0
			} else {
				k -= countMap[count] * count
				delete(countMap, count)
			}
		}
		count++
	}
	fmt.Println("countMap", countMap)

	ans := 0
	for _, v := range countMap {
		ans += v
	}
	return ans

	//for i := 0; i < k; i++ {
	//	key := getKeyOfMinVal(m)
	//	m[key] -= 1
	//	if m[key] == 0 {
	//		delete(m, key)
	//	}
	//}
	//fmt.Println(m)
	//return len(m)
}

func getKeyOfMinVal(m map[int]int) int {
	min, first, key := 0, true, 0
	for k, v := range m {
		if first {
			key, min = k, v
			first = false
		}
		if v < min {
			key, min = k, v
		}
	}
	return key
}
