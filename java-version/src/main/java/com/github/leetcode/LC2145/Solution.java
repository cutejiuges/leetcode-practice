package com.github.leetcode.LC2145;

public class Solution {
    public int numberOfArrays(int[] differences, int lower, int upper) {
        int base = 0;
        int min = 0, max = 0;
        for (int diff : differences) {
            base += diff;
            max = Math.max(max, base);
            min = Math.min(min, base);
            if (max - min > upper - lower) {
                return 0;
            }
        }
        return (upper - lower) - (max - min) + 1;
    }
}
