package com.github.leetcode.LC75;

public class Solution {
    public void sortColors(int[] nums) {
        int p0 = 0, p1 = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            nums[i] = 2;
            if (x <= 1) {
                nums[p1] = 1;
                p1++;
            }
            if (x == 0) {
                nums[p0] = 0;
                p0++;
            }
        }
    }
}
