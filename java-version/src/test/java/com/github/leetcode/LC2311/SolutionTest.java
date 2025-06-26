package com.github.leetcode.LC2311;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void longestSubsequence() {
        Solution solution = new Solution();
        assertEquals(5, solution.longestSubsequence("1001010", 5));
        assertEquals(6, solution.longestSubsequence("00101001", 1));
    }
}
