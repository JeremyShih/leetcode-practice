// 2020/5/10
package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	hasApple := []bool{false, false, true, false, true, true, false}
	fmt.Println(minTime(n, edges, hasApple) == 8)
	hasApple = []bool{false, false, true, false, false, true, false}
	fmt.Println(minTime(n, edges, hasApple) == 6)
	hasApple = []bool{false, false, false, false, false, false, false}
	fmt.Println(minTime(n, edges, hasApple) == 0)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	if ok, _ := Contain(true, hasApple); !ok {
		return 0
	}
	tree := make(map[int][]int)
	for _, e := range edges {
		children := tree[e[0]]
		children = append(children, e[1])
		tree[e[0]] = children
	}
	for k, _ := range tree {
		if childHasApple(tree, k, hasApple) {
			hasApple[k] = true
		}
	}
	count := 0
	for _, v := range hasApple {
		if v {
			count += 2
		}
	}
	count -= 2
	return count
}

func childHasApple(tree map[int][]int, node int, hasApple []bool) bool {
	children := tree[node]
	for _, c := range children {
		if hasApple[c] {
			return true
		}
	}
	for _, c := range children {
		if childHasApple(tree, c, hasApple) {
			return true
		}
	}
	return false
}

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
