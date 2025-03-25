package com.github.leetcode.LC2711;

import java.util.HashSet;
import java.util.Set;

public class Solution {
    public int[][] differenceOfDistinctValues(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[][] ans = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                Set<Integer> bottomRightSet = new HashSet<>();
                int x = i+1, y = j+1;
                while (x < m && y < n) {
                    bottomRightSet.add(grid[x][y]);
                    x++;
                    y++;
                }
                Set<Integer> topLeftSet = new HashSet<>();
                x = i-1;
                y = j-1;
                while (x >= 0 && y >= 0) {
                    topLeftSet.add(grid[x][y]);
                    x--;
                    y--;
                }
                ans[i][j] = Math.abs(bottomRightSet.size() - topLeftSet.size());
            }
        }
        return ans;
    }
}
