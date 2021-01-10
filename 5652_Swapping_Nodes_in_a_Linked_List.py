# 2021/1/10
import time
import math
from fractions import Fraction

# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def swapNodes(self, head: ListNode, k: int) -> ListNode:
        l = []
        cur = head
        while cur != None:
            l.append(cur)
            cur = cur.next
        fromNode = l[k-1]
        toNode = l[-k]
        tmp = toNode.val
        toNode.val = fromNode.val
        fromNode.val = tmp
        return head


start_time = time.time()
slt = Solution()


elapsed_time = time.time() - start_time
print('timp spent:', elapsed_time)
