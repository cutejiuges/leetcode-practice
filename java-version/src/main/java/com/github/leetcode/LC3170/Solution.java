package com.github.leetcode.LC3170;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    public String clearStars(String s) {
        char[] chars = s.toCharArray();
        List<List<Integer>> lists = new ArrayList<>(26);
        for (int i = 0; i < 26; i++) {
            lists.add(new ArrayList<>());
        }
        for (int i = 0; i < chars.length; i++) {
            if (chars[i] != '*') {
                lists.get(chars[i] - 'a').add(i);
            } else {
                for (List<Integer> list : lists) {
                    if (!list.isEmpty()) {
                        int idx = list.remove(list.size() - 1);
                        chars[idx] = '*';
                        break;
                    }
                }
            }
        }

        StringBuilder sb = new StringBuilder();
        for (char ch : chars) {
            sb.append(ch != '*' ? ch : "");
        }
        return sb.toString();
    }
}
