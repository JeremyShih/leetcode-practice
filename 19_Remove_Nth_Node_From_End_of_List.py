# 2019/8/17
# Definition for singly-linked list.
class ListNode:
	def __init__(self, x):
		self.val = x
		self.next = None

class Solution:
	def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
		tmp=head
		l1=self.listNodeToList(head)
		listLen=len(l1)
		if listLen==n:
			head=head.next
		else:
			l1[listLen-n-1].next=l1[listLen-n-1].next.next
		return head
	def listNodeToList(self, head: ListNode) -> list:
		res=[]
		tmp=head
		while tmp:
			res.append(tmp)
			tmp=tmp.next
		return res

slt=Solution()
l1=ListNode(1)
l1.next=ListNode(2)
l1.next.next=ListNode(3)
l1.next.next.next=ListNode(4)
slt.removeNthFromEnd(l1,3)