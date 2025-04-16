package com.github.leetcode.LC2537;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public long countGood(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>();
        long ans = 0;
        int pairs = 0, left = 0;
        for (int num : nums) {
            pairs += cnt.getOrDefault(num, 0);
            cnt.put(num, cnt.getOrDefault(num, 0) + 1);
            while (pairs >= k) {
                pairs -= cnt.getOrDefault(nums[left], 0) - 1;
                cnt.put(nums[left], cnt.getOrDefault(nums[left], 0) - 1);
                left++;
            }
            ans += left;
        }
        return ans;
    }
}
