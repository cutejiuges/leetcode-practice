package com.github.leetcode.LC63;

public class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        if(obstacleGrid == null || obstacleGrid.length == 0 || obstacleGrid[0].length == 0) {
            return 0;
        }
        int m = obstacleGrid.length, n = obstacleGrid[0].length;
        int[][] grid = new int[m][n];

        for(int i = 0; i < m; i++) {
            if(obstacleGrid[i][0] == 1) {
                break;
            }
            grid[i][0] = 1;
        }
        for(int i = 0; i < n; i++) {
            if(obstacleGrid[0][i] == 1) {
                break;
            }
            grid[0][i] = 1;
        }
        for(int i = 1; i < m; i++) {
            for(int j = 1; j < n; j++) {
                if(obstacleGrid[i][j] == 1) {
                    grid[i][j] = 0;
                } else {
                    grid[i][j] = grid[i-1][j] + grid[i][j-1];
                }
            }
        }
        return grid[m-1][n-1];
    }
}
