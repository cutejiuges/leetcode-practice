package com.github.leetcode.LC59;

public class Solution {
    public int[][] generateMatrix(int n) {
        int[][] dir = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
        int curNum = 1, maxNum = n*n;
        int curRow = 0, curCol = 0, curDir = 0;
        int[][] matrix = new int[n][n];
        
        for (; curNum <= maxNum; curNum++) {
            matrix[curRow][curCol] = curNum;
            int newRow = curRow + dir[curDir][0], newCol = curCol + dir[curDir][1];
            if (newRow < 0 || newRow >= n || newCol < 0 || newCol >= n || matrix[newRow][newCol] != 0) {
                curDir = (curDir + 1) % 4;
            }
            curRow = curRow + dir[curDir][0];
            curCol = curCol + dir[curDir][1];
        }
        return matrix;
    }
}
