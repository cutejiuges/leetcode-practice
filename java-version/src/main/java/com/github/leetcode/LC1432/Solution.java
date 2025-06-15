package com.github.leetcode.LC1432;

public class Solution {
    public int maxDiff(int num) {
        String str = String.valueOf(num);
        int max = num, min = num;
        for (int i = 0; i < str.length(); i++) {
            char ch = str.charAt(i);
            if (ch != '9') {
                max = Integer.parseInt(str.replace(ch, '9'));
                break;
            }
        }
        if (str.charAt(0) != '1') {
            min = Integer.parseInt(str.replace(str.charAt(0), '1'));
        } else {
            for (int i = 1; i < str.length(); i++) {
                char ch = str.charAt(i);
                if (ch >= '2') {
                    min = Integer.parseInt(str.replace(ch, '0'));
                    break;
                }
            }
        }
        return max - min;
    }
}
