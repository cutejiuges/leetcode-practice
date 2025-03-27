package com.github.leetcode.LC2712;

public class Solution {
    public long minimumCost(String s) {
        int n = s.length();
        long ans = 0;
        for (int i = 1; i < n; i++) {
            if (s.charAt(i-1) != s.charAt(i)) {
                ans += Math.min(i, n-i);
            }
        }
        return ans;
    }
}
