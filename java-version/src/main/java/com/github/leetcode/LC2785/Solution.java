package com.github.leetcode.LC2785;

import java.util.*;

/**
 * 两次遍历额外存储模拟
 *
 * @author cutejiuge
 * @since 2025/9/11 下午11:10
 */
public class Solution {
    public String sortVowels(String s) {
        Set<Character> vowelSet = new HashSet<>(
            Arrays.asList('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U')
        );
        List<Character> vowels = new ArrayList<>();
        char[] chars = s.toCharArray();
        for (char ch : chars) {
            if (vowelSet.contains(ch)) {
                vowels.add(ch);
            }
        }
        vowels.sort((a, b) -> a - b);
        int idx = 0;
        for (int i = 0; i < chars.length; i++) {
            if (vowelSet.contains(chars[i])) {
                chars[i] = vowels.get(idx++);
            }
        }
        return String.valueOf(chars);
    }
}
