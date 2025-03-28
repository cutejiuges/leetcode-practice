package com.github.leetcode.LC2716;

/**
 * 统计不同的字符数即可
 */

public class Solution {
    public int minimizedStringLength(String s) {
        int mask = 0;
        char[] ss = s.toCharArray();
        for (char ch : ss) {
            mask |= (1 << (ch - 'a'));
        }
        return Integer.bitCount(mask);
    }
}
