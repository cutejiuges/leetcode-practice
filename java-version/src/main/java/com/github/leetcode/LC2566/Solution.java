package com.github.leetcode.LC2566;

public class Solution {
    public int minMaxDifference(int num) {
        String numStr = String.valueOf(num);
        int max = num, min = Integer.MAX_VALUE;
        for (int i = 0; i < numStr.length(); i++) {
            if (numStr.charAt(i) != '9') {
                max = Math.max(max, Integer.parseInt(numStr.replace(numStr.charAt(i), '9')));
                break;
            }
        }
        min = Integer.parseInt(numStr.replace(numStr.charAt(0), '0'));
        return max - min;
    }
}
