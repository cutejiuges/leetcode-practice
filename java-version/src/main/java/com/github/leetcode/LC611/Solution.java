package com.github.leetcode.LC611;

import java.util.Arrays;

/**
 * 排序+查找
 *
 * @author cutejiuge
 * @since 2025/9/26 下午8:50
 */
public class Solution {
    public int triangleNumber(int[] nums) {
        int n = nums.length, ans = 0;
        Arrays.sort(nums);
        for (int i = 0; i < n; i++) {
            int k = i;
            for (int j = i + 1; j < n; j++) {
                // k+1保证最后一个k是正好可以满足条件的
                while (k + 1 < n && nums[k + 1] < nums[i] + nums[j]) {
                    k++;
                }
                ans += Math.max(0, k - j);
            }
        }
        return ans;
    }
}
