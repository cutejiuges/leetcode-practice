package com.github.leetcode.LC2874;

public class Solution {
    /**
     * (nums[i]−nums[j])∗nums[k]尽量大，需要知道j左侧最大值和j右侧最大值，维护前缀最大值和后缀最大值
     * @param nums 数组序列，从中选取3个数字
     * @return 有序序列中 (nums[i] - nums[j]) * nums[k]  可获得的最大值
     */
    public long maximumTripletValue(int[] nums) {
        int n = nums.length;
        int[] sufMax = new int[n+1];
        for (int i = n-1; i > 1; i--) {
            sufMax[i] = Math.max(nums[i], sufMax[i+1]);
        }
        int preMax = nums[0];
        long ans = 0;
        for (int i = 1; i < n-1; i++) {
            ans = Math.max(ans, (long)(preMax - nums[i]) * sufMax[i+1]);
            preMax = Math.max(preMax, nums[i]);
        }
        return ans;
    }
}
