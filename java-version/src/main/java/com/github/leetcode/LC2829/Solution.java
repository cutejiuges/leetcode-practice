package com.github.leetcode.LC2829;

public class Solution {
    public int minimumSum(int n, int k) {
        int leftSum = sumArithmeticSequence(1, 1,  Math.min(n, k/2));
        int rightSum = sumArithmeticSequence(k, 1, n-k/2);
        if (n <= k/2) {
            return leftSum;
        }
        return leftSum + rightSum;
    }

    private int sumArithmeticSequence(int a1, int d, int n) {
        return n*a1 + ((n-1)*n*d / 2);
    }
}
