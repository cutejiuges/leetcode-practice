package com.github.leetcode.LC3904;

public class Solution {
    public int firstStableIndex(int[] nums, int k) {
        int[] suffixMinArr = new int[nums.length];
        suffixMinArr[nums.length - 1] = nums[nums.length - 1];
        for (int i = nums.length - 2; i >= 0; i--) {
            suffixMinArr[i] = Math.min(nums[i], suffixMinArr[i + 1]);
        }

        int prefixMaxVal = nums[0];
        for (int i = 0; i < nums.length; i++) {
            prefixMaxVal = Math.max(prefixMaxVal, nums[i]);
            if (prefixMaxVal - suffixMinArr[i] <= k) return i;
        }
        return -1;
    }
}
