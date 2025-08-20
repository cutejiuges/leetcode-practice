package com.github.leetcode.LC1277;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countSquares() {
        Solution solution = new Solution();
        assertEquals(15, solution.countSquares(new int[][]{{0,1,1,1}, {1,1,1,1}, {0,1,1,1}}));
        assertEquals(7, solution.countSquares(new int[][]{{1,0,1}, {1,1,0}, {1,1,0}}));
    }
}
