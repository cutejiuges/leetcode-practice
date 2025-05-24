package com.github.leetcode.LC3362;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxRemoval() {
        Solution solution = new Solution();
        assertEquals(1, solution.maxRemoval(new int[]{2,0,2}, new int[][]{{0,2}, {0,2}, {1,1}}));
        assertEquals(2, solution.maxRemoval(new int[]{1,1,1,1}, new int[][]{{1,3}, {0,2}, {1,3}, {1,2}}));
        assertEquals(-1, solution.maxRemoval(new int[]{1,2,3,4}, new int[][]{{1,3}}));
    }
}
