// 2020/4/6
package main

import (
	"fmt"
	"math"
)

func main() {
	numCourses := 2
	fmt.Println(numCourses)
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curLevel, maxLevel, maxSum := 1, 1, math.MinInt32
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevelSum, childNumber := 0, len(queue)
		for i := 0; i < childNumber; i++ {
			node := queue[0]
			curLevelSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		if curLevelSum > maxSum {
			maxSum = curLevelSum
			maxLevel = curLevel
		}
		curLevel++
	}
	return maxLevel
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
