package com.github.leetcode.LC3756;

public class Solution {
    private int[] sum, x, cnt, base; // 前缀数组，标记前缀和，前缀数字，非0个数，base的10次幂

    private static final int MAX_N = 100001, MOD = (int) (1e9+7);

    public int[] sumAndMultiply(String s, int[][] queries) {
        this.init(s);
        int[] ans = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int left = queries[i][0], right = queries[i][1] + 1;
            int len = this.cnt[right] - this.cnt[left];
            long xx = this.x[right] - (long) this.x[left] * this.base[len] % MOD + MOD;
            ans[i] = (int) (xx * (this.sum[right] - this.sum[left]) % MOD);
        }
        return ans;
    }

    private void init(String s) {
        this.sum = new int[s.length() + 1];
        this.x = new int[s.length() + 1];
        this.cnt = new int[s.length() + 1];
        this.base = new int[MAX_N];
        // 预构建base数组
        base[0] = 1;
        for (int i = 1; i < MAX_N; i++) {
            base[i] = (int) (base[i - 1] * 10L % MOD);
        }
        // 预构建前缀数组
        for (int i = 0; i < s.length(); i++) {
            int cur = s.charAt(i) - '0';
            this.sum[i + 1] = sum[i] + cur;
            this.x[i + 1] = cur > 0 ? (int) ((this.x[i] * 10L + cur) % MOD) : this.x[i];
            this.cnt[i + 1] = this.cnt[i] + (cur != 0 ? 1 : 0);
        }
    }
}
