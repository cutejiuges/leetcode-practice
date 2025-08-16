package com.github.leetcode.LC1323;

public class Solution {
    public int maximum69Number (int num) {
        char[] ss = String.valueOf(num).toCharArray();
        for (int i = 0; i < ss.length; i++) {
            if (ss[i] == '6') {
                ss[i] = '9';
                break;
            }
        }
        return Integer.parseInt(String.valueOf(ss));
    }
}
