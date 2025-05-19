package com.github.leetcode.LC3024;

import java.util.Arrays;

public class Solution {
    public String triangleType(int[] nums) {
        Arrays.sort(nums);
        int a = nums[0], b = nums[1], c = nums[2];
        if (a + b <= c) {
            return "none";
        }
        if (a == b && b == c) {
            return "equilateral";
        }
        if (a == b || b == c || c == a) {
            return "isosceles";
        }
        return "scalene";
    }
}
