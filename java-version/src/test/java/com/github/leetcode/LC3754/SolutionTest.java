package com.github.leetcode.LC3754;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void sumAndMultiply() {
        Solution solution = new Solution();
        assertEquals(12340L, solution.sumAndMultiply(10203004));
        assertEquals(1L, solution.sumAndMultiply(1000));
    }
}
