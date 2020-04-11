# 2019/12/15
# Definition for singly-linked list.
# class ListNode:
	# def __init__(self, x):
	# 	self.val = x
	# 	self.next = None

class Solution:
	def getDecimalValue(self, head: ListNode) -> int:
		s,a,b="",0,1
		while head:
			s=str(head.val)+s
			head=head.next
		# print(s)
		for i in range(len(s)):
			print(a,b,s[i])
			if int(s[i])>0:
				a+=pow(b,int(s[i]))
			b*=2
		return a

slt=Solution()
