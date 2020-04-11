# 2019/8/13
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        ans = ListNode(0)
        tmp = ans
        while True:
            if l1:
                tmp.val += l1.val
                l1 = l1.next
            if l2:
                tmp.val += l2.val
                l2 = l2.next
            carry = tmp.val // 10
            tmp.val %= 10
            if not l1 and not l2 and carry == 0:
                break
            tmp.next = ListNode(carry)
            tmp = tmp.next
        return ans
    def printListNode(self, l1: ListNode) -> str:
        if not l1:
            return ""
        ans = str(l1.val)
        while l1.next:
            l1 = l1.next
            ans += " -> " + str(l1.val)
        print(ans)
    def arrayToListNode(self, arr1: list) -> ListNode:
        if not arr1:
            return None
        l1 = ListNode(arr1[0])
        l2 = ListNode(arr1[1])
        l2.next = l1
        ans = ListNode(arr1[2])
        ans.next = l2
        return ans


slt=Solution()
# print(slt.printListNode(slt.arrayToListNode([1,2,3])))
l1 = slt.arrayToListNode([9,9,5])
l2 = slt.arrayToListNode([0,0,5])
slt.printListNode(slt.addTwoNumbers(l1,l2))