# 2020/5/24

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def pseudoPalindromicPaths(self, root: TreeNode) -> int:
        l = [root.val]

        def parseTreeNode(path: list, root: TreeNode) -> int:
            count = 0
            if root.left == None and root.right == None:
                result = self.isPesudoPalindromic(path)
                if result:
                    return 1
                else:
                    return 0
            else:
                if root.left != None:
                    count += parseTreeNode(path + [root.left.val], root.left)
                if root.right != None:
                    count += parseTreeNode(path + [root.right.val], root.right)
            return count

        return parseTreeNode(l, root)

    def isPesudoPalindromic(self, l: list) -> bool:
        if len(l) == 1 or len(l) == 0:
            return True

        def CountFrequency(my_list):
            freq = {}
            for item in my_list:
                if item in freq:
                    freq[item] += 1
                else:
                    freq[item] = 1
            return freq

        d = CountFrequency(l)
        if len(l) % 2 == 0:
            for k in d:
                if d[k] % 2 != 0:
                    return False
        else:
            first = True
            for k in d:
                if d[k] % 2 != 0:
                    if not first:
                        return False
                    first = False
        return True
