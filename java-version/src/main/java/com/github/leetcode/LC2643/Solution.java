package com.github.leetcode.LC2643;

public class Solution {
    public int[] rowAndMaximumOnes(int[][] mat) {
        int maxCntOfOnes = -1, index = -1;
        for (int i = 0; i < mat.length; i++) {
            int onesCnt = 0;
            for (int m : mat[i]) {
                onesCnt += m;
            }
            if (onesCnt > maxCntOfOnes) {
                maxCntOfOnes = onesCnt;
                index = i;
            }
        }
        return new int[]{index, maxCntOfOnes};
    }
}
