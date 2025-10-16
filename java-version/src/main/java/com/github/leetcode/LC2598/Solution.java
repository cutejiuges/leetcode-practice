package com.github.leetcode.LC2598;

/**
 * 同余分组
 *
 * @author cutejiuge
 * @since 2025/10/16 下午11:29
 */
public class Solution {
    public int findSmallestInteger(int[] nums, int value) {
        int[] mp = new int[value];
        for (int num : nums) {
            int v = (num % value + value) % value;
            mp[v]++;
        }

        int mex = 0;
        while (mp[mex % value] > 0) {
            mp[mex % value]--;
            mex++;
        }
        return mex;
    }
}
