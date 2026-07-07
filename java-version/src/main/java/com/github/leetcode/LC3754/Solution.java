package com.github.leetcode.LC3754;

public class Solution {
    public long sumAndMultiply(int n) {
        int base = 1;
        long sum = 0, mult = 0;
        for (; n != 0; n /= 10) {
            int remain = n % 10;
            if (remain != 0) {
                sum += remain;
                mult += (long) remain * base;
                base *= 10;
            }
        }
        return sum * mult;
    }
}
