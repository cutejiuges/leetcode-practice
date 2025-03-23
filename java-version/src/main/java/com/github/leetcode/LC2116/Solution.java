package com.github.leetcode.LC2116;

public class Solution {
    public boolean canBeValid(String s, String locked) {
        int n = s.length();
        // 奇数长度字符串不可能是合法的
        if ((n & 1) == 1) {
            return false;
        }

        // 记录未匹配的左括号数量的最值
        int mn = 0, mx = 0;
        for (int i = 0; i < n; i++) {
            if (locked.charAt(i) == '1') { //当前字符不能改
                int cur = s.charAt(i) == '(' ? 1 : -1;
                mx += cur;
                if (mx < 0) { //说明这一段前缀中右括号更多了，不合法
                    return false;
                }
                mn += cur;
            } else { //当前字符可以改动
                mx++; //改成左括号，未匹配的左括号数+1
                mn--; //改成右括号，未匹配的左括号数-1
            }
            mn = Math.max(mn, 0);
        }
        return mn == 0;
    }
}
