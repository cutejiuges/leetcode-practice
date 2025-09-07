package com.github.leetcode.LC1304;

/**
 * 按照相反数依次填入即可
 *
 * @author cutejiuge
 * @since 2025/9/7 下午10:05
 */
public class Solution {
    public int[] sumZero(int n) {
        int[] ans = new int[n];
        int low = 0, high = n - 1;
        int k = n;
        while (low < high) {
            ans[low++] = -k;
            ans[high--] = k;
            k--;
        }
        if ((n & 1) == 1) {
            ans[low] = 0;
        }
        return ans;
    }
}
