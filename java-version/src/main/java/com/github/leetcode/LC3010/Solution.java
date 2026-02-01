package com.github.leetcode.LC3010;

/**
 * 最小代价，维护两个变量标记最小值和次小值
 *
 * @author cutejiuge
 * @since 2026/2/1 下午9:12
 */
public class Solution {
    public int minimumCost(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int first = Integer.MAX_VALUE, second = Integer.MAX_VALUE;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < first) {
                second = first;
                first = nums[i];
            } else if (nums[i] < second) {
                second = nums[i];
            }
        }
        return nums[0] + first + second;
    }
}
