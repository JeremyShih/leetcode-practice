// 2020/5/4
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	m, cur := map[*ListNode]bool{}, head
	for cur != nil {
		if v, _ := m[cur]; v {
			return cur
		}
		m[cur] = true
		cur = cur.Next
	}
	return nil
}
