package com.github.leetcode.LC2270;

public class Solution {
    public int waysToSplitArray(int[] nums) {
        int n = nums.length;
        long left = 0, right = 0;
        for (int num : nums) {
            right += num;
        }

        int ans = 0;
        for (int i = 0; i < n-1; i++) {
            left += nums[i];
            right -= nums[i];
            if (left >= right) {
                ans++;
            }
        }
        return ans;
    }
}
