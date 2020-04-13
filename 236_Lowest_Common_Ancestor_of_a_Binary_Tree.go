// 2020/4/11
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	numbers := []int{3, 5, 1, 6, 2, 0, 8, 7, 4}
	fmt.Println(createTree(numbers))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil && r == nil {
		return nil
	}
	if l != nil {
		return l
	} else {
		return r
	}
}

func createTree(numbers []int) *TreeNode {
	head := TreeNode{}
	cur, i := head, 0
	cur.Val, i = numbers[i], i+1
	fmt.Println(cur.Val, i)
	return &head
}
