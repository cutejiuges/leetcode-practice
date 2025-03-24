package com.github.leetcode.LC2255;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countPrefixes() {
        Solution solution = new Solution();
        assertEquals(3, solution.countPrefixes(new String[]{"a","b","c","ab","bc","abc"}, "abc"));
        assertEquals(2, solution.countPrefixes(new String[]{"a","a"}, "aa"));
    }
}
