package com.github.leetcode.LC2962;

public class Solution {
    public long countSubarrays(int[] nums, int k) {
        int max = Integer.MIN_VALUE;
        for (int num : nums) {
            max = Math.max(max, num);
        }
        int cntMax = 0, left = 0;
        long ans = 0;
        for (int num : nums) {
            if (num == max) {
                cntMax++;
            }
            while (cntMax >= k) {
                if (nums[left] == max) {
                    cntMax--;
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
