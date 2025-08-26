package com.github.leetcode.LC3000;

/**
 * lc3000, 一次遍历
 *
 * @author cutejiuge
 * @since 2025/8/26 上午8:10
 */
public class Solution {
    public int areaOfMaxDiagonal(int[][] dimensions) {
        int maxArea = 0;
        long maxLength = 0L;
        for (int[] dimension : dimensions) {
            int area = dimension[0] * dimension[1];
            long length = (long) dimension[0] * dimension[0] + (long) dimension[1] * dimension[1];
            if (length > maxLength) {
                maxLength = length;
                maxArea = area;
            } else if (length == maxLength) {
                maxArea = Math.max(maxArea, area);
            }
        }
        return maxArea;
    }
}
