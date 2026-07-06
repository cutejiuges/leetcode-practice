package com.github.leetcode.LC1288;

import java.util.Arrays;

public class Solution {
    public int removeCoveredIntervals(int[][] intervals) {
        // 先按照区间起点进行排序,如果起点相同的话长区间在前
        Arrays.sort(intervals, (o1, o2) -> {
            if (o1[0] != o2[0]) {
                return o1[0] - o2[0];
            }
            return o2[1] - o1[1];
        });
        // 排序之后进行遍历，只要右节点没被覆盖就更新右节点，答案加一
        int ans = 1, right = intervals[0][1];
        for (int i = 1; i < intervals.length; i++) {
            if (intervals[i][1] > right) {
                ans++;
                right = intervals[i][1];
            }
        }
        return ans;
    }
}
