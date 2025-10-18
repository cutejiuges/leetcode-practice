package com.github.leetcode.LC3397;

import java.util.Arrays;

/**
 * 排序后进行贪心比对
 *
 * @author cutejiuge
 * @since 2025/10/18 下午10:16
 */
public class Solution {
    public int maxDistinctElements(int[] nums, int k) {
        int pre = Integer.MIN_VALUE;
        Arrays.sort(nums);
        // 从小元素开始进行比对
        int diff = 0;
        for (int num : nums) {
            int v = Math.min(Math.max(pre + 1, num - k), num + k);
            if (v > pre) {
                diff++;
                pre = v;
            }
        }
        return diff;
    }
}
