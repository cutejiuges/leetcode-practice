package com.github.leetcode.LC1706;

public class Solution {
    public int[] findBall(int[][] grid) {
        int len = grid[0].length;
        int[] ans = new int[len];
        for(int i = 0; i < len; i++) {
            int curCol = i;
            for(int[] row : grid) {
                int dir = row[curCol], nextCol = curCol + dir;
                if(nextCol < 0 || nextCol >= len || row[nextCol] != dir) {
                    curCol = -1;
                    break;
                }
                curCol = nextCol;
            }
            ans[i] = curCol;
        }
        return ans;
    }
}
