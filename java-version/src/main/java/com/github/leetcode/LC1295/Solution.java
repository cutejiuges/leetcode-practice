package com.github.leetcode.LC1295;

public class Solution {
    public int findNumbers(int[] nums) {
        int ans = 0;
        for (int num : nums) {
            ans += (this.count10Bits(num) & 1) == 0 ? 1 : 0;
        }
        return ans;
    }

    private int count10Bits(int num) {
        int x = num, ans = 0;
        while (x > 0) {
            x /= 10;
            ans++;
        }
        return ans;
    }
}
