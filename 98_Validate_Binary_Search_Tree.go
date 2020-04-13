// 2020/4/11
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	s := "3+2*2"
	fmt.Println(s == "7")
	fmt.Println(math.Inf(1), math.Inf(-1))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}
	return low < root.Val && root.Val < high && helper(root.Left, low, root.Val) && helper(root.Right, root.Val, high)
}
