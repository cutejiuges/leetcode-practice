package com.github.leetcode.LC2799;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class Solution {
    public int countCompleteSubarrays(int[] nums) {
        int distinct = this.differentNums(nums);
        int left = 0, ans = 0;
        Map<Integer, Integer> map = new HashMap<>(distinct);
        for (int right = 0; right < nums.length; right++) {
            map.merge(nums[right], 1, Integer::sum);
            while (map.size() >= distinct) {
                map.merge(nums[left], -1, Integer::sum);
                if (map.get(nums[left]) == 0) {
                    map.remove(nums[left]);
                }
                left++;
            }
            ans += left;
        }
        return ans;
    }

    private int differentNums(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for (int num : nums) {
            set.add(num);
        }
        return set.size();
    }
}
