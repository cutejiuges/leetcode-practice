package com.github.leetcode.LC3442;

public class Solution {
    public int maxDifference(String s) {
        int[] cnt = new int[26];
        for (int i = 0; i < s.length(); i++) {
            cnt[s.charAt(i) - 'a']++;
        }
        int max = Integer.MIN_VALUE, min = Integer.MAX_VALUE;
        for (int num : cnt) {
            if (num == 0) {
                continue;
            }
            if ((num & 1) == 1) {
                max = Math.max(max, num);
            } else {
                min = Math.min(min, num);
            }
        }
        return max - min;
    }
}
