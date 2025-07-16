package com.github.leetcode.LC3201;

public class Solution {
    public int maximumLength(int[] nums) {
        int ans = 0;
        int[][] patterns = new int[][]{{0, 0}, {0, 1}, {1, 0}, {1,1}};
        for (int[] pattern : patterns) {
            int cnt = 0;
            for (int num : nums) {
                if ((num & 1) == pattern[cnt & 1]) {
                    cnt++;
                }
            }
            ans = Math.max(ans, cnt);
        }
        return ans;
    }
}
