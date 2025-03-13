package com.github.leetcode.LC3306;

import java.util.HashMap;

public class Solution {
    public long countOfSubstrings(String word, int k) {
        char[] ss = word.toCharArray();
        return calculate(ss, k) - calculate(ss, k+1);
    }

    private long calculate(char[] ss, int k) {
        long ans = 0;
        HashMap<Character, Integer> vowelCntMap = new HashMap<>();
        int consonantCnt = 0;
        int low = 0;
        for (int i = 0; i < ss.length; i++) {
            if ("aeiou".indexOf(ss[i]) >= 0) {
                vowelCntMap.put(ss[i], vowelCntMap.getOrDefault(ss[i], 0) + 1);
            } else {
                consonantCnt++;
            }

            while (vowelCntMap.size() >= 5 && consonantCnt >= k) {
                char leftChar = ss[low];
                if (vowelCntMap.containsKey(leftChar)) {
                    vowelCntMap.put(leftChar, vowelCntMap.getOrDefault(leftChar, 0) - 1);
                    if (vowelCntMap.get(leftChar) == 0) {
                        vowelCntMap.remove(leftChar);
                    }
                } else {
                    consonantCnt--;
                }
                low++;
            }
            ans += low;
        }
        return ans;
    }
}
