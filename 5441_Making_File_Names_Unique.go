// 2020/6/21
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	names := []string{"pes", "fifa", "gta", "pes(2019)"}
	ans := []string{"pes", "fifa", "gta", "pes(2019)"}
	//fmt.Println(reflect.DeepEqual(getFolderNames(names), ans))
	//names = []string{"gta", "gta(1)", "gta", "avalon"}
	//ans = []string{"gta", "gta(1)", "gta(2)", "avalon"}
	//fmt.Println(reflect.DeepEqual(getFolderNames(names), ans))
	names = []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}
	ans = []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"}
	fmt.Println(reflect.DeepEqual(getFolderNames(names), ans))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func getFolderNames(names []string) []string {
	m := make(map[string]int)
	for i := 0; i < len(names); i++ {
		v := names[i]
		fmt.Println(i, v)
		if cou, ok := m[v]; ok {

			fmt.Println(cou, ok)
			new := v + "(" + strconv.Itoa(cou) + ")"
			m[v] += 1
			m[new] = 1
			names[i] = new
		} else {
			m[v] = 1
		}
	}
	fmt.Println(names)
	return names
}
