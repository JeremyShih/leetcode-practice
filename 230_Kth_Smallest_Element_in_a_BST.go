// 2020/4/11 failed to contruct a binary search tree for testing
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	root := []int{3, 1, 4, math.MinInt32, 2}
	// k := 1
	fmt.Println(construct(root))
	// fmt.Println(kthSmallest(construct(root), k) == 1)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	stack = []int{}
	bst(root)
	return stack[k-1]
}

func bst(root *TreeNode) {
	if root.root == nil {
		return
	}
	bst(root.root.Left)
	stack = append(stack, root.root.Val)
	bst(root.root.Right)
}

func construct(input []int) TreeNode {
	var head TreeNode
	for _, n := range input {
		head.insert(n)
	}
	fmt.Println(head)
	return head
}

var stack []int
