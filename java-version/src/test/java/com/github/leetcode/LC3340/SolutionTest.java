package com.github.leetcode.LC3340;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isBalanced() {
        Solution solution = new Solution();
        assertFalse(solution.isBalanced("1234"));
        assertTrue(solution.isBalanced("24123"));
    }
}
