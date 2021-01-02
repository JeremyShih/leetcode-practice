// 2020/8/23
package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	n, rounds := 4, []int{1, 3, 1, 2}
	fmt.Println(reflect.DeepEqual(mostVisited(n, rounds), []int{1, 2}))
	n, rounds = 2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}
	fmt.Println(reflect.DeepEqual(mostVisited(n, rounds), []int{2}))
	n, rounds = 7, []int{1, 3, 5, 7}
	fmt.Println(reflect.DeepEqual(mostVisited(n, rounds), []int{1, 2, 3, 4, 5, 6, 7}))
	n, rounds = 3, []int{3, 2, 1, 2, 1, 3, 2, 1, 2, 1, 3, 2, 3, 1}
	fmt.Println(reflect.DeepEqual(mostVisited(n, rounds), []int{1, 3}))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func mostVisited(n int, rounds []int) []int {
	rec := make(map[int]int)
	countMax := 0
	for i := 1; i <= n; i++ {
		rec[i] = 0
	}
	rec[rounds[0]] += 1
	for i := 1; i < len(rounds); i++ {
		cur := rounds[i-1] + 1
		if cur > n {
			cur = cur % n
		}
		for cur != rounds[i] {
			rec[cur]++
			cur++
			if cur > n {
				cur = cur % n
			}
		}
		rec[cur]++
		countMax = rec[cur]
	}
	//fmt.Println(countMax, rec)
	res := []int{}
	for k, v := range rec {
		if v == countMax {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	//fmt.Println(res)
	return res
}
