package com.github.leetcode.LC2506;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int similarPairs(String[] words) {
        Map<Integer, Integer> map = new HashMap<>();
        int ans = 0;
        for (String ss : words) {
            int mask = 0;
            for (int i = 0; i < ss.length(); i++) {
                mask |= 1 << (ss.charAt(i) - 'a');
            }
            int cnt = map.getOrDefault(mask, 0);
            ans += cnt;
            map.put(mask, cnt+1);
        }
        return ans;
    }
}
