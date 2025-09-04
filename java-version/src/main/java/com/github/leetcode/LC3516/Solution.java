package com.github.leetcode.LC3516;

/**
 * 求绝对值距离就行，签到题
 *
 * @author cutejiuge
 * @since 2025/9/4 下午11:14
 */
public class Solution {
    public int findClosest(int x, int y, int z) {
        int a = Math.abs(x - z), b = Math.abs(y - z);
        if (a == b) return 0;
        return a < b ? 1 : 2;
    }
}
