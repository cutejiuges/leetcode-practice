package com.github.leetcode.LC3355;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isZeroArray() {
        Solution solution = new Solution();
        assertTrue(solution.isZeroArray(new int[]{1, 0, 1}, new int[][]{{0, 2}}));
        assertFalse(solution.isZeroArray(new int[]{4,3,2,1}, new int[][]{{1,3}, {0,2}}));
    }
}
