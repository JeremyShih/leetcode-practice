// 2020/6/13
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	head := &TreeNode{}
	//fmt.Println(rob(head) == 0)
	//head = &TreeNode{3, nil, nil}
	//fmt.Println(rob(head) == 3)
	head = &TreeNode{3, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{3, nil, &TreeNode{1, nil, nil}}}
	fmt.Println(rob(head) == 7)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo map[*TreeNode]int

func robMemo(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if v, ok := memo[node]; ok {
		return v
	}

	grandChildren := 0
	if node.Left != nil {
		grandChildren += robMemo(node.Left.Left) + robMemo(node.Left.Right)
	}
	if node.Right != nil {
		grandChildren += robMemo(node.Right.Left) + robMemo(node.Right.Right)
	}

	memo[node] = node.Val + grandChildren
	child := robMemo(node.Left) + robMemo(node.Right)
	if child > memo[node] {
		memo[node] = child
	}
	return memo[node]
}

func rob(root *TreeNode) int {
	memo = make(map[*TreeNode]int)
	return robMemo(root)
}
