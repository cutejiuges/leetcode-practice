package com.github.leetcode.LC1266;

/**
 * 遍历模拟即可
 *
 * @author cutejiuge
 * @since 2026/1/12 下午8:47
 */
public class Solution {
    public int minTimeToVisitAllPoints(int[][] points) {
        int ans = 0;
        for (int i = 1; i < points.length; i++) {
            int x = Math.abs(points[i][0] - points[i-1][0]), y = Math.abs(points[i][1] - points[i-1][1]);
            ans += Math.max(x, y);
        }
        return ans;
    }
}
