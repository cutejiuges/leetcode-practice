package com.github.leetcode.LC1493;

/**
 * LC1493 转换题意即为找到最长的值包含1个0的子串，滑动窗口
 *
 * @author cutejiuge
 * @since 2025/8/24 上午6:59
 */
public class Solution {
    public int longestSubarray(int[] nums) {
        int cnt0 = 0, left = 0, ans = 0;
        for (int right = 0; right < nums.length; right++) {
            cnt0 += nums[right] == 0 ? 1 : 0;
            while (cnt0 > 1) {
                cnt0 -= nums[left] == 0 ? 1 : 0;
                left++;
            }
            ans = Math.max(ans, right - left);
        }
        return ans;
    }
}
