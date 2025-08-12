package com.github.leetcode.LC2787;

public class Solution {
    public int numberOfWays(int n, int x) {
        int mod = (int) 1e9+7;
        long[] dp = new long[n + 1];
        dp[0] = 1;
        for (int i = 1; quickPower(i, x) <= n; i++) {
            int val = quickPower(i, x);
            for (int j = n; j >= val; j--) {
                dp[j] = (dp[j] + dp[j-val]) % mod;
            }
        }
        return (int) dp[n];
    }

    private int quickPower(int a, int b) {
        int ans = 1;
        while (b != 0) {
            if ((b & 1) == 1) {
                ans *= a;
            }
            a *= a;
            b >>= 1;
        }
        return ans;
    }
}
