// 2020/4/19
package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"time"
	"util"
)

func main() {
	start := time.Now()
	orders := [][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}
	output := [][]string{{"Table", "Beef Burrito", "Ceviche", "Fried Chicken", "Water"}, {"3", "0", "2", "1", "0"}, {"5", "0", "1", "0", "1"}, {"10", "1", "0", "0", "0"}}
	fmt.Println(reflect.DeepEqual(displayTable(orders), output))
	orders = [][]string{{"James", "12", "Fried Chicken"}, {"Ratesh", "12", "Fried Chicken"}, {"Amadeus", "12", "Fried Chicken"}, {"Adam", "1", "Canadian Waffles"}, {"Brianna", "1", "Canadian Waffles"}}
	output = [][]string{{"Table", "Canadian Waffles", "Fried Chicken"}, {"1", "2", "0"}, {"12", "0", "3"}}
	fmt.Println(reflect.DeepEqual(displayTable(orders), output))
	orders = [][]string{{"Laura", "2", "Bean Burrito"}, {"Jhon", "2", "Beef Burrito"}, {"Melissa", "2", "Soda"}}
	output = [][]string{{"Table", "Bean Burrito", "Beef Burrito", "Soda"}, {"2", "1", "1", "1"}}
	fmt.Println(reflect.DeepEqual(displayTable(orders), output))
	elapsed := time.Now().Sub(start)
	fmt.Println("time spent:", elapsed)
}

type Table struct {
	Number int
	Orders map[string]int
}

func displayTable(orders [][]string) [][]string {
	tables, manus, tableNumbers := make(map[int]Table), []string{}, []int{}
	for _, o := range orders {
		i, _ := strconv.Atoi(o[1])
		if ok, _ := util.Contain(o[2], manus); !ok {
			manus = append(manus, o[2])
		}
		if ok, _ := util.Contain(i, tableNumbers); !ok {
			tableNumbers = append(tableNumbers, i)
		}
		if table, ok := tables[i]; ok {
			table.Orders[o[2]] += 1
		} else {
			table := Table{Number: i, Orders: make(map[string]int)}
			table.Orders[o[2]] = 1
			tables[i] = table
		}
	}
	sort.Ints(tableNumbers)
	sort.Strings(manus)
	res, tmp := [][]string{}, append([]string{"Table"}, manus...)
	res = append(res, tmp)
	for _, v := range tableNumbers {
		tmp = []string{strconv.Itoa(v)}
		table := tables[v]
		for _, v := range manus {
			if count, ok := table.Orders[v]; ok {
				tmp = append(tmp, strconv.Itoa(count))
			} else {
				tmp = append(tmp, "0")
			}
		}
		res = append(res, tmp)
	}
	return res
}
