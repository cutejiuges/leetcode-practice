package com.github.leetcode.LC3136;

public class Solution {
    public boolean isValid(String word) {
        if (word == null || word.length() < 3) {
            return false;
        }
        boolean containsVowel = false, containsConsonant = false;
        char[] chars = word.toCharArray();
        for (char ch : chars) {
            if (Character.isLetter(ch)) {
                char c = Character.toLowerCase(ch);
                if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                    containsVowel = true;
                } else {
                    containsConsonant = true;
                }
            } else if (!Character.isDigit(ch)) {
                return false;
            }
        }
        return containsVowel && containsConsonant;
    }
}
