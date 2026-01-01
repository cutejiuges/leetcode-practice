package com.github.leetcode.LC66;

/**
 * 逆序遍历进行模拟即可
 *
 * @author cutejiuge
 * @since 2026/1/1 下午10:08
 */
public class Solution {
    public int[] plusOne(int[] digits) {
        for (int i = digits.length - 1; i >= 0; i--) {
            digits[i] = (digits[i] + 1) % 10;
            if (digits[i] != 0) {
                return digits;
            }
        }
        int[] ans = new int[digits.length + 1];
        ans[0] = 1;
        return ans;
    }
}
