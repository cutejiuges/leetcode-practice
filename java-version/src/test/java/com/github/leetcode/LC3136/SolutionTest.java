package com.github.leetcode.LC3136;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isValid() {
        Solution solution = new Solution();
        assertTrue(solution.isValid("234Adas"));
        assertFalse(solution.isValid("b3"));
        assertFalse(solution.isValid("a3$e"));
    }
}
