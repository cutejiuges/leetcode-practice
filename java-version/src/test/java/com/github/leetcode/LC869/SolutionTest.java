package com.github.leetcode.LC869;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void reorderedPowerOf2() {
        Solution solution = new Solution();
        assertTrue(solution.reorderedPowerOf2(1));
        assertFalse(solution.reorderedPowerOf2(10));
    }
}
