package com.github.leetcode.LC1007;

public class Solution {
    private final int INF = 0x3f3f3f;
    public int minDominoRotations(int[] tops, int[] bottoms) {
        int ans = Math.min(rotation(tops, bottoms, tops[0]), rotation(tops, bottoms, bottoms[0]));
        return ans == this.INF ? -1 : ans;
    }

    private int rotation(int[] tops, int[] bottoms, int target) {
        int topCnt = 0, bottomCnt = 0;
        for (int i = 0; i < tops.length; i++) {
            if (tops[i] != target && bottoms[i] != target) {
                return this.INF;
            }
            if (tops[i] != target) {
                topCnt++;
            } else if (bottoms[i] != target) {
                bottomCnt++;
            }
        }
        return Math.min(topCnt, bottomCnt);
    }
}
