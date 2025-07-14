from typing import Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    @staticmethod
    def getDecimalValue(head: Optional[ListNode]) -> int:
        node = head
        ans = 0
        while node:
            ans *= 2
            ans += node.val
            node = node.next
        return ans
