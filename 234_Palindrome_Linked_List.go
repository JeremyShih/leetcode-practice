// 2020/4/19
package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main() {
	start := time.Now()
	s1 := []int{2, 4}
	fmt.Println(!isPalindrome(sliceToListNode(s1)))
	s1 = []int{2, 1}
	fmt.Println(!isPalindrome(sliceToListNode(s1)))
	s1 = []int{2, 1, 1, 2}
	fmt.Println(isPalindrome(sliceToListNode(s1)))
	s1 = []int{3, 2, 1, 1, 2, 3}
	fmt.Println(isPalindrome(sliceToListNode(s1)))
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func isPalindrome(head *ListNode) bool {
	s := listNodeToSlice(head)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
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
