package com.github.leetcode.LC1290;

// Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

public class Solution {
    public int getDecimalValue(ListNode head) {
        ListNode node = head;
        int ans = 0;
        while (node != null) {
            ans *= 2;
            ans += node.val;
            node = node.next;
        }
        return ans;
    }
}
