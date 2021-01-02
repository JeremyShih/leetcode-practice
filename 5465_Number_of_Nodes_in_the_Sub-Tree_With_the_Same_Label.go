// 2020/7/19
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
	labels := "abaedcd"
	fmt.Println(reflect.DeepEqual(countSubTrees(n, edges, labels), []int{2, 1, 1, 1, 1, 1, 1}))
	n = 4
	edges = [][]int{{0, 1}, {1, 2}, {0, 3}}
	labels = "bbbb"
	fmt.Println(reflect.DeepEqual(countSubTrees(n, edges, labels), []int{4, 2, 1, 1}))
	n = 5
	edges = [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}
	labels = "aabab"
	fmt.Println(reflect.DeepEqual(countSubTrees(n, edges, labels), []int{3, 2, 1, 1, 1}))
	n = 6
	edges = [][]int{{0, 1}, {0, 2}, {1, 3}, {3, 4}, {4, 5}}
	labels = "cbabaa"
	fmt.Println(reflect.DeepEqual(countSubTrees(n, edges, labels), []int{1, 2, 1, 1, 2, 1}))
	n = 7
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	labels = "aaabaaa"
	fmt.Println(reflect.DeepEqual(countSubTrees(n, edges, labels), []int{6, 5, 4, 1, 3, 2, 1}))
	n = 4
	edges = [][]int{{0, 2}, {0, 3}, {1, 2}}
	labels = "aeed"
	fmt.Println(reflect.DeepEqual(countSubTrees(n, edges, labels), []int{1, 1, 2, 1}))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

var ans []int
var tree map[int][]int
var label string

func countSubTrees(n int, edges [][]int, labels string) []int {
	ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}
	tree = make(map[int][]int)
	tree[0] = []int{}
	label = labels
	for len(edges) > 0 {
		e := edges[0]
		if _, ok := tree[e[0]]; ok {
			tree[e[0]] = append(tree[e[0]], e[1])
			tree[e[1]] = []int{}
		} else if _, ok := tree[e[1]]; ok {
			tree[e[1]] = append(tree[e[1]], e[0])
			tree[e[0]] = []int{}
		} else {
			edges = append(edges, e)
		}
		edges = edges[1:]
	}
	for i := 0; i < n; i++ {
		for _, n := range tree[i] {
			dfs(i, n)
		}
	}
	return ans
}

func dfs(root int, node int) {
	if label[root] == label[node] {
		ans[root] += 1
	}
	for _, n := range tree[node] {
		if label[root] == label[n] {
			ans[root] += ans[n]
		} else {
			dfs(root, n)
		}
	}
}
