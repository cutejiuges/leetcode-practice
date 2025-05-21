package com.github.leetcode.LC3356;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minZeroArray() {
        Solution solution = new Solution();
        assertEquals(2, solution.minZeroArray(new int[]{2,0,2}, new int[][]{{0,2,1}, {0,2,1}, {1,1,3}}));
        assertEquals(-1, solution.minZeroArray(new int[]{4,3,2,1}, new int[][]{{1,3,2}, {0,2,1}}));
    }
}
