package com.github.leetcode.LC3340;

public class Solution {
    public boolean isBalanced(String num) {
        int cnt = 0;
        for (int i = 0; i < num.length(); i++) {
            if ((i & 1) == 0) {
                cnt += num.charAt(i) - '0';
            } else {
                cnt -= num.charAt(i) - '0';
            }
        }
        return cnt == 0;
    }
}
