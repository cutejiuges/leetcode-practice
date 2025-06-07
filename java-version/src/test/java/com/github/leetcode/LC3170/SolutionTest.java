package com.github.leetcode.LC3170;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void clearStars() {
        Solution solution = new Solution();
        assertEquals("aab", solution.clearStars("aaba*"));
        assertEquals("abc", solution.clearStars("abc"));
        assertEquals("", solution.clearStars("d*o*"));
    }
}
