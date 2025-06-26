package com.github.leetcode.LC2311;

public class Solution {
    public int longestSubsequence(String s, int k) {
        int n = s.length(), m = 32 - Integer.numberOfLeadingZeros(k);
        if (n < m) {
            return n;
        }
        int ans = m;
        int sufVal = Integer.parseInt(s.substring(n - m), 2);
        if (sufVal > k) {
            ans--;
        }
        for (int i = 0; i < n - m; i++) {
            ans += '1' - s.charAt(i);
        }
        return ans;
    }
}
