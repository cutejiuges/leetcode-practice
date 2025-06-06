package com.github.leetcode.LC2434;

import java.util.ArrayDeque;
import java.util.Deque;

public class Solution {
    public String robotWithString(String s) {
        // 计算某个位置的后缀最小值
        int len = s.length();
        char[] sufMin = new char[len + 1];
        sufMin[len] = Character.MAX_VALUE;
        for (int i = len - 1; i >= 0; i--) {
            sufMin[i] = (char) Math.min(sufMin[i+1], s.charAt(i));
        }

        // 从最小后缀贪心构建答案
        StringBuilder ans = new StringBuilder();
        Deque<Character> stack = new ArrayDeque<>();
        for (int i = 0; i < len; i++) {
            stack.push(s.charAt(i));
            while (!stack.isEmpty() && stack.peek() <= sufMin[i + 1]) {
                ans.append(stack.pop());
            }
        }
        return ans.toString();
     }
}
