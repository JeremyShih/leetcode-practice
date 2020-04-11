# 2019/10/13
"""
# Definition for a Node.
class Node:
	def __init__(self, val, left, right, next):
		self.val = val
		self.left = left
		self.right = right
		self.next = next
"""
class Solution:
	def connect(self, root: 'Node') -> 'Node':
		if root:
			self.connect(root.left)
			self.connect(root.right)
			self.connectNext(root.left,root.right)
		return root
	def connectNext(self, left, right):
		while left:
			left.next=right
			left=left.right
			right=right.left

slt=Solution()

