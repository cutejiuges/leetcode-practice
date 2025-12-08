package com.github.leetcode.LC1925;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countTriples() {
        Solution solution = new Solution();
        assertEquals(2, solution.countTriples(5));
        assertEquals(4, solution.countTriples(10));
    }
}
