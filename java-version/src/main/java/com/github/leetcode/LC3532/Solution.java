package com.github.leetcode.LC3532;

public class Solution {
    public boolean[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        int[] collect = new int[n];
        for (int i = 1; i < n; i++) {
            collect[i] = nums[i] - nums[i - 1] > maxDiff ? i : collect[i - 1];
        }
        boolean[] ans = new boolean[queries.length];
        for (int i = 0; i < queries.length; i++) {
            ans[i] = collect[queries[i][0]] == collect[queries[i][1]];
        }
        return ans;
    }
}
