package com.github.leetcode.LC2595;

public class Solution {
    public int[] evenOddBit(int n) {
        int[] ans = new int[2];
        int pos = 0;
        while(n > 0) {
            ans[pos] += n & 1;
            n >>= 1;
            pos ^= 1;
        }
        return ans;
    }
}
