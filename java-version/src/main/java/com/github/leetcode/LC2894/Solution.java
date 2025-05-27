package com.github.leetcode.LC2894;

public class Solution {
    public int differenceOfSums(int n, int m) {
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            if (i % m != 0) {
                ans += i;
            } else {
                ans -= i;
            }
        }
        return ans;
    }
}
