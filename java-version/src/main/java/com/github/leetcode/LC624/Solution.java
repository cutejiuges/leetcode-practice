package com.github.leetcode.LC624;

import java.util.List;

public class Solution {
    public int maxDistance(List<List<Integer>> arrays) {
        int ans = 0;
        int min = Integer.MAX_VALUE / 2, max = Integer.MIN_VALUE / 2;
        for (List<Integer> list : arrays) {
            int first = list.get(0), last = list.get(list.size() - 1);
            ans = Math.max(ans, Math.max(max - first, last - min));
            min = Math.min(min, first);
            max = Math.max(max, last);
        }
        return ans;
    }
}