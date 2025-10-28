package com.github.leetcode.LC3354;

/**
 * 前缀和分界判断即可
 *
 * @author cutejiuge
 * @since 2025/10/28 下午11:22
 */
public class Solution {
    public int countValidSelections(int[] nums) {
        int sum = 0, cur = 0, ans = 0;
        for (int num : nums) {
            sum += num;
        }
        for (int num : nums) {
            cur += num;
            if (num == 0 && cur * 2 == sum) {
                ans += 2;
            }
            if (num == 0 && Math.abs(cur * 2 - sum) == 1) {
                ans += 1;
            }
        }
        return ans;
    }
}
