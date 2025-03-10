package com.github.leetcode.LC2269;

public class Solution {
    public int divisorSubstrings(int num, int k) {
        String ss = String.valueOf(num);
        int len = ss.length(), ans = 0;
        for (int i = 0; i <= len-k; i++) {
            int t = Integer.parseInt(ss.substring(i, i+k));
            if (t != 0 && num % t == 0) {
                ans++;
            }
        }
        return ans;
    }
}
