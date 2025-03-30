package com.github.leetcode.LC2109;

public class Solution {
    public String addSpaces(String s, int[] spaces) {
        StringBuilder sb = new StringBuilder();
        int length = s.length(), spaceIdx = 0;
        for (int i = 0; i < length; i++) {
            if (spaceIdx < spaces.length && i == spaces[spaceIdx]) {
                sb.append(' ');
                spaceIdx++;
            }
            sb.append(s.charAt(i));
        }
        return sb.toString();
    }
}
