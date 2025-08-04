package com.github.leetcode.LC904;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int totalFruit(int[] fruits) {
        Map<Integer, Integer> map = new HashMap<>();
        int left = 0, ans = 0;
        for (int i = 0; i < fruits.length; i++) {
            map.merge(fruits[i], 1, Integer::sum);
            while (map.size() > 2) {
                map.merge(fruits[left], -1, Integer::sum);
                if (map.get(fruits[left]) == 0) {
                    map.remove(fruits[left]);
                }
                left++;
            }
            ans = Math.max(ans, i - left + 1);
        }
        return ans;
    }
}
