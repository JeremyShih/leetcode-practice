// 2020/5/3
package main

import (
	"fmt"
	// "math"
	"time"
)

func main() {
	start := time.Now()
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destCity(paths) == "Sao Paulo")
	paths = [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	fmt.Println(destCity(paths) == "A")
	paths = [][]string{{"A", "Z"}}
	fmt.Println(destCity(paths) == "Z")
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func destCity(paths [][]string) string {
	m := map[string]int{}
	for _, p := range paths {
		m[p[0]] = 0
		if v, ok := m[p[1]]; ok && v == 0 {
			m[p[1]] = 0
		} else {
			m[p[1]] = 1
		}
	}
	// fmt.Println(m)
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return ""
}
