package com.github.leetcode.LC3169;

import java.util.Arrays;

public class Solution {
    public int countDays(int days, int[][] meetings) {
        int ans = days;
        Arrays.sort(meetings, (a, b) -> a[0] - b[0]);
        int start = meetings[0][0], end = meetings[0][1];
        for (int i = 1; i < meetings.length; i++) {
            if (meetings[i][0] > end) {
                ans -= end - start + 1;
                start = meetings[i][0];
            }
            end = Math.max(end, meetings[i][1]);
        }
        ans -= end - start + 1;
        return ans;
    }
}
