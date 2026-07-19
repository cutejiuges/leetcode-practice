package com.github.leetcode.LC1081;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void smallestSubsequence() {
        Solution solution = new Solution();
        assertEquals("abc", solution.smallestSubsequence("bcabc"));
        assertEquals("acdb", solution.smallestSubsequence("cbacdcbc"));
    }
}
