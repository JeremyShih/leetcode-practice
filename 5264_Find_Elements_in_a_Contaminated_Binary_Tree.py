# 2019/11/17
# Definition for a binary tree node.
# class TreeNode:
# 	def __init__(self, x):
# 		self.val = x
# 		self.left = None
# 		self.right = None

class FindElements:

	def __init__(self, root: TreeNode):
		self.dic={}
		def recover(root, val):
			self.dic[val]=root
			root.val=val
			if root.left:
				recover(root.left, val*2+1)
			if root.right:
				recover(root.right, val*2+2)
		recover(root, 0)
		return

	def find(self, target: int) -> bool:
		if target in self.dic:
			return True
		else:
			return False


# Your FindElements object will be instantiated and called as such:
# obj = FindElements(root)
# param_1 = obj.find(target)