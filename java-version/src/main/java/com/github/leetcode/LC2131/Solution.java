package com.github.leetcode.LC2131;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class Solution {
    public int longestPalindrome(String[] words) {
        Map<String, Integer> map = new HashMap<>();
        for (String word : words) {
            map.put(word, map.getOrDefault(word, 0) + 1);
        }
        Set<String> set = new HashSet<>();

        int ans = 0;
        boolean odd = false;
        for (String word : map.keySet()) {
            int cnt = map.get(word);
            String reverseWord = new StringBuilder(word).reverse().toString();
            if (word.equals(reverseWord)) {
                if ((cnt & 1) == 1) {
                    odd = true;
                }
                ans += (cnt / 2 * 2) * 2;
            } else {
                if (!set.contains(word)) {
                    set.add(word);
                    set.add(reverseWord);
                    ans += Math.min(map.getOrDefault(word, 0), map.getOrDefault(reverseWord, 0)) * 2 * 2;
                }
            }
        }
        if (odd) {
            ans += 2;
        }
        return ans;
    }
}
