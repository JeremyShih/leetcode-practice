# 2019/11/3
class Solution:
	def sortList(self, head: ListNode) -> ListNode:
		if not head:
			return head
		l=[]
		while head:
			l.append(head.val)
			head=head.next
		l.sort()
		head=cur=ListNode(l[0])
		for i in range(1,len(l)):
			cur.next=ListNode(l[i])
			cur=cur.next
		return head

# Definition for singly-linked list.
class ListNode:
	def __init__(self, x):
		self.val = x
		self.next = None