package com.github.leetcode.LC2140;

public class Solution {
    public long mostPoints(int[][] questions) {
        int len = questions.length;
        long[] dp = new long[len+1];
        for (int i = len-1; i >= 0; i--) {
            dp[i] = Math.max(dp[i+1], questions[i][0] + dp[Math.min(len, i+questions[i][1]+1)]);
        }
        return dp[0];
    }
}
