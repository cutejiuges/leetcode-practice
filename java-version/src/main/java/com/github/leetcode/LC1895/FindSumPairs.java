package com.github.leetcode.LC1895;

import java.util.HashMap;
import java.util.Map;

public class FindSumPairs {
    private int[] nums1, nums2;
    private Map<Integer, Integer> map;

    public FindSumPairs(int[] nums1, int[] nums2) {
        this.map = new HashMap<>();
        for (int num : nums2) {
            map.merge(num, 1, Integer::sum);
        }
        this.nums1 = nums1;
        this.nums2 = nums2;
    }

    public void add(int index, int val) {
        int old = this.nums2[index];
        this.map.merge(old, -1, Integer::sum);
        this.nums2[index] += val;
        this.map.merge(this.nums2[index], 1, Integer::sum);
    }

    public int count(int tot) {
        int ans = 0;
        for (int num : this.nums1) {
            ans += this.map.getOrDefault(tot - num, 0);
        }
        return ans;
    }
}
