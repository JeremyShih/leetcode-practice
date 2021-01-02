// 2020/4/23
package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main() {
	start := time.Now()
	n, pos := []int{3, 2, 0, -4}, 1
	fmt.Println(hasCycle(sliceToListNode(n, pos)))
	n, pos = []int{1, 2}, 0
	fmt.Println(hasCycle(sliceToListNode(n, pos)))
	n, pos = []int{1, 2}, -1
	fmt.Println(!hasCycle(sliceToListNode(n, pos)))
	n, pos = []int{1}, -1
	fmt.Println(!hasCycle(sliceToListNode(n, pos)))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m, cur := map[*ListNode]bool{}, head
	for cur != nil {
		if v, _ := m[cur]; v {
			return true
		}
		m[cur] = true
		cur = cur.Next
	}
	return false
}

func sliceToListNode(input []int, index int) *ListNode {
	head, m := &ListNode{Val: input[0]}, []*ListNode{}
	cur := head
	m = append(m, cur)
	input = input[1:]
	for len(input) > 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		m = append(m, cur)
		cur.Val = input[0]
		input = input[1:]
	}
	if index > -1 {
		cur.Next = m[index]
	}
	return head
}
