package com.github.leetcode.LC2918;

public class Solution {
    public long minSum(int[] nums1, int[] nums2) {
        long[] res1 = calculate(nums1), res2 = calculate(nums2);
        if (res1[0] > res2[0] && res2[1] == 0) {
            return -1;
        }
        if (res2[0] > res1[0] && res1[1] == 0) {
            return -1;
        }
        return Math.max(res1[0], res2[0]);
    }

    private long[] calculate(int[] nums) {
        long sum = 0, cnt0 = 0;
        for (int num : nums) {
            sum += num == 0 ? 1 : num;
            cnt0 += num == 0 ? 1 : 0;
        }
        return new long[]{sum, cnt0};
    }
}
