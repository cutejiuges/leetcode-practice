package com.github.leetcode.LC342;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isPowerOfFour() {
        Solution solution = new Solution();
        assertTrue(solution.isPowerOfFour(16));
        assertTrue(solution.isPowerOfFour(1));
        assertFalse(solution.isPowerOfFour(5));
    }
}
