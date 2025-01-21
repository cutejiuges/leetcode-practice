package com.github.LC2218;

import java.util.List;

public class Solution {
    public int maxValueOfCoins(List<List<Integer>> piles, int k) {
        int n = piles.size();
        int[][] dp = new int[n+1][k+1]; //dp[i][j] 表示考虑从 [0,i] 堆物品中挑选，背包容量为 j 时的最大价值

        for(int i = 0; i < n; i++) { //枚举第i堆
            List<Integer> pile = piles.get(i);
            for(int j = 1; j <= k; j++) { //枚举面临第j次选择
                //不选择这一组的任何数字
                dp[i+1][j] = dp[i][j];
                int val = 0;
                for(int w = 1; w <= Math.min(j, pile.size()); w++) { //枚举，尝试从当前堆选择w个
                    val += pile.get(w-1);
                    dp[i+1][j] = Math.max(dp[i+1][j], dp[i][j-w]+val);
                }
            }
        }
        return dp[n][k];
    }
}
