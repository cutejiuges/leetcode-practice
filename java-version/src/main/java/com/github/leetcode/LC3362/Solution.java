package com.github.leetcode.LC3362;

import java.util.Arrays;
import java.util.PriorityQueue;

public class Solution {
    public int maxRemoval(int[] nums, int[][] queries) {
        Arrays.sort(queries, (a, b) -> a[0] - b[0]);
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> b - a);
        int[] diffArray = new int[nums.length+1];
        int sumDiff = 0, j = 0;
        for (int i = 0; i < nums.length; i++) {
            sumDiff += diffArray[i];
            while (j < queries.length && queries[j][0] <= i) {
                pq.add(queries[j][1]);
                j++;
            }
            while (sumDiff < nums[i] && !pq.isEmpty() && pq.peek() >= i) {
                sumDiff++;
                diffArray[pq.poll()+1]--;
            }
            if (sumDiff < nums[i]) {
                return -1;
            }
        }
        return pq.size();
    }
}
