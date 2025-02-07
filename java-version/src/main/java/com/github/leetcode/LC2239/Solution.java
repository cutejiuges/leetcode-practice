package com.github.leetcode.LC2239;

public class Solution {
    public int findClosestNumber(int[] nums) {
        int ans = Integer.MAX_VALUE;
        for(int num : nums) {
            if (Math.abs(num) < Math.abs(ans)) {
                ans = num;
            } else if(Math.abs(ans) == Math.abs(num)) {
                ans = Math.max(num, ans);
            }
        }
        return ans;
    }
}
