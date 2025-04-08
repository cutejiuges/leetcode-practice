package com.github.leetcode.LC3396;

import java.util.HashSet;
import java.util.Set;

public class Solution {
    // 直接通过遍历，倒序检查
    public int minimumOperations(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int i = nums.length - 1; i >= 0; i--) {
            if (set.contains(nums[i])) {
                return i / 3 + 1;
            }
            set.add(nums[i]);
        }
        return 0;
    }
}
