package com.github.leetcode.LC2566;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minMaxDifference() {
        Solution solution = new Solution();
        assertEquals(99009, solution.minMaxDifference(11891));
        assertEquals(99, solution.minMaxDifference(90));
    }
}
