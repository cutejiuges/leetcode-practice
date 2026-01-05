package com.github.leetcode.LC1975;

/**
 * 按照负数的个数进行判断
 *
 * @author cutejiuge
 * @since 2026/1/5 下午11:20
 */
public class Solution {
    public long maxMatrixSum(int[][] matrix) {
        long absSum = 0;
        int cnt = 0, minAbs = Integer.MAX_VALUE;
        for (int[] row : matrix) {
            for (int num : row) {
                int curAbs = Math.abs(num);
                minAbs = Math.min(minAbs, curAbs);
                absSum += curAbs;
                if (num < 0) {
                    cnt++;
                }
            }
        }
        return (cnt & 1) == 0 ? absSum : absSum - 2L * minAbs;
    }
}
