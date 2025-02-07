package com.github.leetcode.LC3066;

import java.util.PriorityQueue;

public class Solution {
    public int minOperations(int[] nums, int k) {
        PriorityQueue<Long> pq = new PriorityQueue<>();
        for(long num : nums) {
            pq.offer(num);
        }

        int ans = 0;
        while(pq.peek() < k) {
            long x = pq.poll(), y = pq.poll();
            pq.offer(2*x+y);
            ans++;
        }
        return ans;
    }
}
