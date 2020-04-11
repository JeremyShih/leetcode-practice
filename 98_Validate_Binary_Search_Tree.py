# 2019/9/3
# Definition for a binary tree node.
class TreeNode:
	def __init__(self, x):
		self.val = x
		self.left = None
		self.right = None

class Solution:
	def isValidBST(self, root: TreeNode) -> bool:
		def valid(node, low, high):
			if not node: return True
			return low < node.val < high and valid(node.left, low, node.val) and valid(node.right, node.val, high)
		return valid(root, float('-inf'), float('inf'))

slt=Solution()