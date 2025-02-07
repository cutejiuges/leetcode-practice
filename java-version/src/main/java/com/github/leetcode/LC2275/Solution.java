package com.github.leetcode.LC2275;

public class Solution {
    public int largestCombination(int[] candidates) {
        //求出数组中的最大值
        int mx = 0;
        for (int num : candidates) {
            mx = Math.max(num, mx);
        }
        //计算最大值的二进制位数
        int maxLen = 32 - Integer.numberOfLeadingZeros(mx);

        //枚举非零二进制位，计算需要多少个数字
        int ans = 0;
        for (int i = 0; i < maxLen; i++) {
            int curCnt = 0;
            for (int num : candidates) {
                curCnt += (num >> i & 1);
            }
            ans = Math.max(ans, curCnt);
        }
        return ans;
    }
}
