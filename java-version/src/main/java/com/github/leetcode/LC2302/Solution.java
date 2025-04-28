package com.github.leetcode.LC2302;

public class Solution {
    public long countSubarrays(int[] nums, long k) {
        long ans = 0, score = 0;
        int left = 0;
        for (int i = 0; i < nums.length; i++) {
            score += nums[i];
            while (score * (i - left + 1) >= k) {
                score -= nums[left++];
            }
            ans += i - left + 1;
        }
        return ans;
    }
}
