package com.github.leetcode.LC3195;

/**
 * LC3195日题练习，直接遍历
 *
 * @author cutejiuge
 * @since 2025/8/22 下午7:59
 */
public class Solution {
    public int minimumArea(int[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        int m = grid.length, n = grid[0].length;
        int minRow = m, minCol = n, maxRow = 0, maxCol = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    minRow = Math.min(minRow, i);
                    minCol = Math.min(minCol, j);
                    maxRow = Math.max(maxRow, i);
                    maxCol = Math.max(maxCol, j);
                }
            }
        }
        return (maxRow - minRow + 1) * (maxCol - minCol + 1);
    }
}
