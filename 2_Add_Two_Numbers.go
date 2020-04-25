// 2020/4/19
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	start := time.Now()
	s1, s2 := []int{2, 4, 3}, []int{5, 6, 4}
	ans := []int{7, 0, 8}
	fmt.Println(reflect.DeepEqual(listNodeToSlice(addTwoNumbers(sliceToListNode(s1), sliceToListNode(s2))), ans))
	s1, s2 = []int{5, 4, 3}, []int{5, 6, 4}
	ans = []int{0, 1, 8}
	fmt.Println(reflect.DeepEqual(listNodeToSlice(addTwoNumbers(sliceToListNode(s1), sliceToListNode(s2))), ans))
	s1, s2 = []int{5, 4, 3, 3}, []int{5, 6, 4, 2}
	ans = []int{0, 1, 8, 5}
	//fmt.Println(listNodeToSlice(addTwoNumbers(sliceToListNode(s1), sliceToListNode(s2))))
	fmt.Println(reflect.DeepEqual(listNodeToSlice(addTwoNumbers(sliceToListNode(s1), sliceToListNode(s2))), ans))
	s1 = []int{1, 2, 3, 4}
	fmt.Println(s1, reflect.DeepEqual(s1, listNodeToSlice(sliceToListNode(s1))))
	s1 = []int{1, 2, 3, 4, 5}
	fmt.Println(s1, reflect.DeepEqual(s1, listNodeToSlice(sliceToListNode(s1))))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, carry := &ListNode{}, 0
	cur := head
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Val = carry % 10
		carry /= 10
		if l1 != nil || l2 != nil || carry > 0 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToListNode(input []int) *ListNode {
	head := &ListNode{Val: input[0]}
	cur := head
	input = input[1:]
	for len(input) > 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = input[0]
		input = input[1:]
	}
	return head
}

func listNodeToSlice(l *ListNode) []int {
	s := []int{}
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return s
}
