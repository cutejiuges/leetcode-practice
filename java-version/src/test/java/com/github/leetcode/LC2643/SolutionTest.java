package com.github.leetcode.LC2643;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void rowAndMaximumOnes() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{0,1}, solution.rowAndMaximumOnes(new int[][]{{0,1}, {1,0}}));
        assertArrayEquals(new int[]{1,2}, solution.rowAndMaximumOnes(new int[][]{{0,0,0}, {0,1,1}}));
        assertArrayEquals(new int[]{1,2}, solution.rowAndMaximumOnes(new int[][]{{0,0}, {1,1}, {0,0}}));
    }
}
