// 2020/7/19
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles + 1
	for numBottles >= numExchange {
		exchange := (numBottles - (numBottles % numExchange)) / numExchange
		ans += exchange
		numBottles = numBottles%numExchange + exchange
	}
	return ans
}
