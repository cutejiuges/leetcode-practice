package com.github.leetcode.LC781;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public int numRabbits(int[] answers) {
        Map<Integer, Integer> map = new HashMap<>();
        int ans = 0;
        for (int answer : answers) {
            if (map.getOrDefault(answer, 0) == 0) {
                ans += answer + 1;
                map.put(answer, answer);
            } else {
                map.merge(answer, -1, Integer::sum);
            }
        }
        return ans;
    }
}
