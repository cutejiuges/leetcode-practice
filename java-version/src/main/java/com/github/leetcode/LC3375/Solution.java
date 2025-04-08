package com.github.leetcode.LC3375;

import java.util.HashSet;
import java.util.Set;

public class Solution {
    // 本质上是比较不同数字的个数
    public int minOperations(int[] nums, int k) {
        Set<Integer> set = new HashSet<>();
        for (int m : nums) {
            if (m < k) {
                return -1;
            } else if (m > k) {
                set.add(m);
            }
        }
        return set.size();
    }
}
