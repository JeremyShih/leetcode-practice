// 2020/7/26
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func countPairs(root *TreeNode, distance int) int {
	pool := []*TreeNode{root}
	ans := 0
	pairs := []string{}
	for len(pool) > 0 {
		l := len(pool)
		for i := 0; i < l; i++ {
			n := pool[0]
			dist, nodes := distanceToLeaves(n)
			//fmt.Println(dist, nodes, pairs)
			for j := 0; j < len(nodes)-1; j++ {
				for k := j + 1; k < len(nodes); k++ {
					if y, _ := Contain(strconv.Itoa(nodes[j])+"_"+strconv.Itoa(nodes[k]), pairs); y {
						continue
					}
					if y, _ := Contain(strconv.Itoa(nodes[k])+"_"+strconv.Itoa(nodes[j]), pairs); y {
						continue
					}
					if dist[nodes[j]]+dist[nodes[k]] <= distance {
						//fmt.Println(nodes[j], nodes[k])
						ans += 1
						pairs = append(pairs, strconv.Itoa(nodes[j])+"_"+strconv.Itoa(nodes[k]))
						pairs = append(pairs, strconv.Itoa(nodes[k])+"_"+strconv.Itoa(nodes[j]))
					}
				}
			}

			if n.Left != nil {
				pool = append(pool, n.Left)
			}
			if n.Right != nil {
				pool = append(pool, n.Right)
			}
			pool = pool[1:]
		}
	}
	fmt.Println(len(pairs), pairs)
	return ans
}

func distanceToLeaves(root *TreeNode) (map[int]int, []int) {
	pool := []*TreeNode{root}
	rec := map[*TreeNode]int{root: 0}
	leavesMap := make(map[int]int)
	leaves := []int{}
	for len(pool) > 0 {
		l := len(pool)
		for i := 0; i < l; i++ {
			n := pool[0]
			if n.Right != nil {
				pool = append(pool, n.Right)
				rec[n.Right] = rec[n] + 1
			}
			if n.Left != nil {
				pool = append(pool, n.Left)
				rec[n.Left] = rec[n] + 1
			}
			if n.Right == nil && n.Left == nil {
				leavesMap[n.Val] = rec[n]
				leaves = append(leaves, n.Val)
			}
			pool = pool[1:]
		}
	}
	return leavesMap, leaves
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
