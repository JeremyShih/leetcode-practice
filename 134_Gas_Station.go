package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas...) < sum(cost...) {
		return -1
	}
	ans, tank := 0, 0
	for i, v := range gas {
		left := v - cost[i]
		if tank+left < 0 {
			ans = i + 1
			tank = 0
		} else {
			tank += left
		}
	}
	return ans
}

func sum(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost) == 3)
	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost) == -1)

}
