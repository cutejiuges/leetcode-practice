package com.github.leetcode.LC3633;

public class Solution {
    public int earliestFinishTime(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {
        int landWater = solve(landStartTime, landDuration, waterStartTime, waterDuration);
        int waterLand = solve(waterStartTime, waterDuration, landStartTime, landDuration);
        return Math.min(landWater, waterLand);
    }

    private int solve(int[] firstStartTime, int[] firstDuration, int[] secondStartTime, int[] secondDuration) {
        int minFinish = Integer.MAX_VALUE;
        for (int i = 0; i < firstStartTime.length; i++) {
            minFinish = Math.min(minFinish, firstStartTime[i] + firstDuration[i]);
        }
        int res = Integer.MAX_VALUE;
        for (int i = 0; i < secondStartTime.length; i++) {
            res = Math.min(res, Math.max(secondStartTime[i], minFinish) + secondDuration[i]);
        }
        return res;
    }
}
