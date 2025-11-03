package com.github.leetcode.LC1578;

/**
 * 贪心 + 栈
 *
 * @author cutejiuge
 * @since 2025/11/3 下午11:08
 */
public class Solution {
    public int minCost(String colors, int[] neededTime) {
        char preColor = colors.charAt(0);
        int preTime = neededTime[0], cost = 0;
        for (int i = 1; i < neededTime.length; i++) {
            if (preColor == colors.charAt(i)) {
                cost += Math.min(preTime, neededTime[i]);
                preTime = Math.max(preTime, neededTime[i]);
            } else {
                preColor = colors.charAt(i);
                preTime = neededTime[i];
            }
        }
        return cost;
    }
}
