package com.github.leetcode.LC3356;

import java.util.Arrays;

public class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        int queryLength = queries.length;
        int low = 0, high = queryLength + 1;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (this.check(nums, queries, mid)) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        return low <= queryLength ? low : -1;
    }

    private boolean check(int[] nums, int[][] queries, int k) {
        int[] diffArray = new int[nums.length];
        int[][] subQueries = Arrays.copyOfRange(queries, 0, k);
        for (int[] query : subQueries) {
            diffArray[query[0]] += query[2];
            if (query[1] < nums.length - 1) {
                diffArray[query[1]+1] -= query[2];
            }
        }

        int sumDiff = 0;
        for (int i = 0; i < nums.length; i++) {
            sumDiff += diffArray[i];
            if (nums[i] > sumDiff) {
                return false;
            }
        }
        return true;
    }
}
