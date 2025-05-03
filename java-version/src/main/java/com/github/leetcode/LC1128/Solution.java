package com.github.leetcode.LC1128;

public class Solution {
    public int numEquivDominoPairs(int[][] dominoes) {
        int[][] visited = new int[10][10];
        int ans = 0;
        for (int[] domino : dominoes) {
            int x = Math.min(domino[0], domino[1]), y = Math.max(domino[0], domino[1]);
            ans += visited[x][y];
            visited[x][y]++;
        }
        return ans;
    }
}
