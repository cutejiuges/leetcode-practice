package com.github.leetcode.LC1780;

public class Solution {
    public boolean checkPowersOfThree(int n) {
        return powerSumOfX(n, 3);
    }

    private boolean powerSumOfX(int num, int x) {
        if (num < 1) {
            return false;
        }
        while (num != 0) {
            if (num % x > 1) {
                return false;
            }
            num /= x;
        }
        return true;
    }
}
