package com.github.leetcode.LC2278;

public class Solution {
    public int percentageLetter(String s, char letter) {
        int cnt = 0, len = s.length();
        for (int i = 0; i < len; i++) {
            if (s.charAt(i) == letter) {
                cnt++;
            }
        }
        return cnt * 100 / len;
    }
}
