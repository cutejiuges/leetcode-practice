package com.github.leetcode.LC3304;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public char kthCharacter(int k) {
        List<Character> ss = new ArrayList<>();
        ss.add('a');
        while (k > ss.size()) {
            List<Character> cur = new ArrayList<>(ss);
            for (char ch : cur) {
                char c = (char)('a' + (ch - 'a' + 1) % 26);
                ss.add(c);
            }
        }
        return ss.get(k - 1);
    }
}
