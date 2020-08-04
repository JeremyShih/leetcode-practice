# 2020/6/13
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def rob(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        def robRecursive(node):
            if not node:
                return 0

            grandchildren = (
                robRecursive(node.left.left) +
                robRecursive(node.left.right) if node.left else 0
            ) + (
                robRecursive(node.right.left) +
                robRecursive(node.right.right) if node.right else 0
            )

            return max(
                node.val + grandchildren,
                robRecursive(node.left) + robRecursive(node.right))

        def robMemoization(node):
            if not node:
                return 0
            elif node in memoize:
                return memoize[node]

            grandchildren = (
                robMemoization(node.left.left) +
                robMemoization(node.left.right) if node.left else 0
            ) + (
                robMemoization(node.right.left) +
                robMemoization(node.right.right) if node.right else 0
            )

            memoize[node] = max(
                node.val + grandchildren,
                robMemoization(node.left) + robMemoization(node.right))
            return memoize[node]

        memoize = dict()
        return robMemoization(root)
