package com.github.leetcode.LC326;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isPowerOfThree() {
        Solution solution = new Solution();
        assertTrue(solution.isPowerOfThree(27));
        assertFalse(solution.isPowerOfThree(0));
        assertTrue(solution.isPowerOfThree(9));
        assertFalse(solution.isPowerOfThree(45));
    }
}
