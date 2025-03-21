package com.github.leetcode.LC2680;

public class Solution {
    public long maximumOr(int[] nums, int k) {
        int allOr = 0, fixed = 0;
        for (int m : nums) {
            fixed |= allOr & m;
            allOr |= m;
        }

        long ans = 0;
        for (int m : nums) {
            ans = Math.max(ans, (allOr ^ m) | fixed | (long)m << k);
        }
        return ans;
    }
}
