// 2020/5/17
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	startTime := []int{1, 2, 3}
	endTime := []int{3, 2, 7}
	queryTime := 4
	fmt.Println(busyStudent(startTime, endTime, queryTime) == 1)
	startTime = []int{4}
	endTime = []int{4}
	queryTime = 4
	fmt.Println(busyStudent(startTime, endTime, queryTime) == 1)
	startTime = []int{4}
	endTime = []int{4}
	queryTime = 5
	fmt.Println(busyStudent(startTime, endTime, queryTime) == 0)
	startTime = []int{1, 1, 1, 1}
	endTime = []int{1, 3, 2, 4}
	queryTime = 7
	fmt.Println(busyStudent(startTime, endTime, queryTime) == 0)
	startTime = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	endTime = []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	queryTime = 5
	fmt.Println(busyStudent(startTime, endTime, queryTime) == 5)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
