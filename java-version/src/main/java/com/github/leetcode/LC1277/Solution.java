package com.github.leetcode.LC1277;

public class Solution {
    public int countSquares(int[][] matrix) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
            return 0;
        }
        int m = matrix.length, n = matrix[0].length;
        int[][] dp = new int[m][n];
        int ans = 0;
        // 枚举当前的右下角
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                // 如果是靠边的位置，只能取当前一个位置
                if (i == 0 || j == 0) {
                    dp[i][j] = matrix[i][j];
                } else if (matrix[i][j] == 0) { // 如果右下角元素是0,这点的正方形个数一定是0
                    dp[i][j] = 0;
                } else { // 如果右下角元素是1,从三个状态中进行转移
                    dp[i][j] = Math.min(Math.min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1;
                }
                ans += dp[i][j];
            }
        }
        return ans;
    }
}
