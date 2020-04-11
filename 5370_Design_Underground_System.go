// 2020/3/29
package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.CheckIn(45, "Leyton", 3)
	obj.CheckOut(45, "Waterloo", 15)
	fmt.Println(obj.GetAverageTime(startStation, endStation) == 12)
}

type UndergroundSystem struct {
	check map[string]int
}

func Constructor() UndergroundSystem {

}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {

}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {

}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {

}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
