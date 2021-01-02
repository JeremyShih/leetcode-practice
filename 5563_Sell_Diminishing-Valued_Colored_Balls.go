// 2020/11/8
package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	inventory, orders := []int{2, 5}, 4
	fmt.Println(maxProfit(inventory, orders) == 14)
	inventory, orders = []int{3, 5}, 6
	fmt.Println(maxProfit(inventory, orders) == 19)
	inventory, orders = []int{2, 8, 4, 10, 6}, 20
	fmt.Println(maxProfit(inventory, orders) == 110)
	inventory, orders = []int{1000000000}, 1000000000
	fmt.Println(maxProfit(inventory, orders) == 21)
	inventory, orders = []int{76}, 22
	fmt.Println(maxProfit(inventory, orders) == 1441)
	inventory, orders = []int{773160767}, 252264991
	//fmt.Println(maxProfit(inventory, orders) == 70267492)
	fmt.Println(maxProfit(inventory, orders))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxProfit(inventory []int, orders int) int {
	mod := math.Pow(10, 9) + 7
	val := math.Mod(0, mod)
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	//fmt.Println(inventory)
	l := len(inventory)

	for orders > 0 {
		i := 0
		//fmt.Println(inventory[i], inventory[i+1], inventory[i] < inventory[i+1])
		for i+1 < l && inventory[i] <= inventory[i+1] {
			i++
		}
		//fmt.Println(i)
		target := inventory[i] - 1
		if i+1 < l {
			target = inventory[i+1]
		}
		if l == 1 {
			target = inventory[i] - orders
		}
		fmt.Println("ori", inventory[i], target)
		fmt.Println(float64(inventory[i])+float64(target)+1, float64(inventory[i])-float64(target))
		//diff := (float64(inventory[i]) + float64(target) + 1) * (float64(inventory[i]) - float64(target)) / 2
		//diff = big.NewInt(inventory[i]+target+1) * big.NewInt(inventory[i]-target)
		mul := new(big.Float).Mul(big.NewFloat(float64(inventory[i]+target+1)), big.NewFloat(float64(inventory[i]-target)))
		div := new(big.Float).Quo(mul, big.NewFloat(2))
		mqwepu:=div%big.NewFloat(3)
		diff, _ := div.Float64()
		fmt.Println(fmt.Sprintf("%f", diff), fmt.Sprintf("%f", math.Mod(diff, mod)))
		val = math.Mod(val+diff, mod)
		fmt.Println(inventory[i], target, diff)
		orders -= inventory[i] - target
		inventory[i] = target

		//val = math.Mod(val+float64(inventory[i]), mod)
		//inventory[i] -= 1
		//orders--
		fmt.Println(inventory, orders, val)
	}

	return int(val)
}
