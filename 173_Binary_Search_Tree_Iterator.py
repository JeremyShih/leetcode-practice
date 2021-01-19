# 2021/1/10
import time
import math
from fractions import Fraction


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class BSTIterator:
    def __init__(self, root: TreeNode):
        self.nodes_sorted = []
        self.index = -1
        self._inorder(root)

    def _inorder(self, root: TreeNode):
        return

    def next(self) -> int:
        return 0

    def hasNext(self) -> bool:
        return True


        # Your BSTIterator object will be instantiated and called as such:
        # obj = BSTIterator(root)
        # param_1 = obj.next()
        # param_2 = obj.hasNext()
start_time = time.time()


elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
