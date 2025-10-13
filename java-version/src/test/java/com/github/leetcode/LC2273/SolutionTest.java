package com.github.leetcode.LC2273;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void removeAnagrams() {
        Solution solution = new Solution();
        assertArrayEquals(new String[]{"abba","cd"}, solution.removeAnagrams(new String[]{"abba","baba","bbaa","cd","cd"}).toArray());
        assertArrayEquals(new String[]{"a","b","c","d","e"}, solution.removeAnagrams(new String[]{"a","b","c","d","e"}).toArray());
    }
}
