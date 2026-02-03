package com.github.leetcode.LC3637;

/**
 * 一次循环判断转弯总数
 *
 * @author cutejiuge
 * @since 2026/2/3 下午11:09
 */
public class Solution {
    public boolean isTrionic(int[] nums) {
        // 起始递增
        if (nums[0] >= nums[1]) {
            return false;
        }
        int cnt = 1;
        // 从后续位置进行遍历来判断
        for (int i = 1; i < nums.length - 1; i++) {
            // 如果相邻元素有相等，不满足
            if (nums[i] == nums[i - 1] || nums[i] == nums[i + 1]) {
                return false;
            }
            // 相连的3个元素有拐弯，递增一个
            if (nums[i] > nums[i - 1] && nums[i] > nums[i + 1] ||
                nums[i] < nums[i - 1] && nums[i] < nums[i + 1]) {
                cnt++;
            }
        }
        return cnt == 3;
    }
}
