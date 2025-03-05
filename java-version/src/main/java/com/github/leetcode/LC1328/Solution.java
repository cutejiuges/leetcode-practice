package com.github.leetcode.LC1328;

public class Solution {
    public String breakPalindrome(String palindrome) {
        int n = palindrome.length();
        if (n <= 1) {
            return "";
        }

        char[] ss = palindrome.toCharArray();
        for (int i = 0; i < n / 2; i++) {
            if (ss[i] != 'a') {
                ss[i] = 'a';
                return String.valueOf(ss);
            }
        }
        ss[n-1] = 'b';
        return String.valueOf(ss);
    }
}
