package com.github.leetcode.LC812;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void largestTriangleArea() {
        Solution solution = new Solution();
        assertEquals(2.00000, solution.largestTriangleArea(new int[][]{{0,0}, {0,1}, {1,0}, {0,2},{2,0}}), 1e-5);
        assertEquals(0.50000, solution.largestTriangleArea(new int[][]{{1,0}, {0,0}, {0,1}}));
    }
}
