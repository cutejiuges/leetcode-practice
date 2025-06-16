package com.github.leetcode.LC2016;

public class Solution {
    public int maximumDifference(int[] nums) {
        int ans = -1, preMin = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (preMin >= nums[i]) {
                preMin = nums[i];
            } else {
                ans = Math.max(ans, nums[i] - preMin);
            }
        }
        return ans;
    }
}
