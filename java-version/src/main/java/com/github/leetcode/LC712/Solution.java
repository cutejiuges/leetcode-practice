package com.github.leetcode.LC712;

/**
 * 动态规划最长公共子串
 *
 * @author cutejiuge
 * @since 2026/1/10 下午7:40
 */
public class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        char[] ss1 = s1.toCharArray(), ss2 = s2.toCharArray();
        // 先计算保留全部字符的和
        int total = 0;
        for (char c : ss1) {
            total += c;
        }
        for (char c : ss2) {
            total += c;
        }

        int m = ss1.length, n = ss2.length;
        // 动态规划，计算保留的字符的和
        int[][] dp = new int[m + 1][n + 1];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                // 如果当前字符相等，保留的和加上当前字符
                if (ss1[i] == ss2[j]) {
                    dp[i+1][j+1] = dp[i][j] + ss1[i];
                } else { // 如果当前字符不相等，取前一段字符中比较大的那个
                    dp[i+1][j+1] = Math.max(dp[i][j+1], dp[i+1][j]);
                }
            }
        }
        // 从total中把保留下来的减掉就是不要的，因为计算的时候存了一份，但是实际上每个字符串都会存一份，所以需要减掉两倍
        return total - dp[m][n]*2;
    }
}
