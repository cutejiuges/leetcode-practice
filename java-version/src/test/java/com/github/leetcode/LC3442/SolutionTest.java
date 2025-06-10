package com.github.leetcode.LC3442;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxDifference() {
        Solution solution = new Solution();
        assertEquals(3, solution.maxDifference("aaaaabbc"));
        assertEquals(1, solution.maxDifference("abcabcab"));
    }
}
