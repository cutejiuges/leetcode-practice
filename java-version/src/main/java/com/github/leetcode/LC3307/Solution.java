package com.github.leetcode.LC3307;

public class Solution {
    public char kthCharacter(long k, int[] operations) {
        int ans = 0, t = 0;
        while (k != 1) {
            t = 63 - Long.numberOfLeadingZeros(k);
            if ((1L << t) == k) {
                t--;
            }
            k = k - (1L << t);
            if (operations[t] != 0) {
                ans++;
            }
        }
        return (char) ('a' + (ans % 26));
    }
}
