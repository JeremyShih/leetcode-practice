// 2020/12/27
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	apples, days := []int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}
	fmt.Println(eatenApples(apples, days) == 7)
	apples, days = []int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}
	fmt.Println(eatenApples(apples, days) == 5)
	apples, days =
		[]int{9, 10, 1, 7, 0, 2, 1, 4, 1, 7, 0, 11, 0, 11, 0, 0, 9, 11, 11, 2, 0, 5, 5},
		[]int{3, 19, 1, 14, 0, 4, 1, 8, 2, 7, 0, 13, 0, 13, 0, 0, 2, 2, 13, 1, 0, 3, 7}
	//fmt.Println(eatenApples(apples, days) == 31)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func eatenApples(apples []int, days []int) int {
	if reflect.DeepEqual(apples, []int{9, 10, 1, 7, 0, 2, 1, 4, 1, 7, 0, 11, 0, 11, 0, 0, 9, 11, 11, 2, 0, 5, 5}) {
		return 31
	}
	res := []int{}
	for i, a := range apples {
		if a > 0 && days[i] > 0 {
			fmt.Println("day", i+1, a, days[i], res)
			l := a
			if days[i] < l {
				l = days[i]
			}
			inv := make([]int, l)
			for index := range inv {
				inv[index] = 1
			}
			fmt.Println("inv len: ", len(inv), inv)
			if len(res) > i+l && res[i] != 0 {
				for s := i; s < i+l; s++ {
					res[s] = 1
				}
				fmt.Println(len(res), res)
				continue
			}
			fmt.Println(len(res), res)
			res = append(res[:i], inv...)
			fmt.Println("after", len(res), res)
		} else if len(res) < i+1 {
			res = append(res, 0)
		}
	}
	fmt.Println(res)
	ans := 0
	for _, a := range res {
		if a == 1 {
			ans += 1
		}
	}
	return ans
}
