package com.github.leetcode.LC2141;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxRunTime() {
        Solution solution = new Solution();
        assertEquals(4, solution.maxRunTime(2, new int[]{3, 3, 3}));
        assertEquals(2, solution.maxRunTime(2, new int[]{1, 1, 1, 1}));
    }
}
