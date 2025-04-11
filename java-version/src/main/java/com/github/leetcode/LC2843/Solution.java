package com.github.leetcode.LC2843;

public class Solution {
    public int countSymmetricIntegers(int low, int high) {
        int ans = 0;
        for (int i = low; i <= high; i++) {
            String str = Integer.toString(i);
            if ((str.length() & 1) > 0) {
                continue;
            }
            int judge = 0;
            for (int j = 0; j < str.length() / 2; j++) {
                judge += str.charAt(j) - '0';
            }
            for (int j = str.length() / 2; j < str.length(); j++) {
                judge -= str.charAt(j) - '0';
            }
            ans += judge == 0 ? 1 : 0;
        }
        return ans;
    }
}
