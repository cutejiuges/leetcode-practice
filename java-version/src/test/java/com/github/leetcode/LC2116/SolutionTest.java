package com.github.leetcode.LC2116;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void canBeValid() {
        Solution solution = new Solution();
        assertTrue(solution.canBeValid("))()))", "010100"));
        assertTrue(solution.canBeValid("()()", "0000"));
        assertFalse(solution.canBeValid(")", "0"));
        assertTrue(solution.canBeValid("(((())(((())", "111111010111"));
    }
}
