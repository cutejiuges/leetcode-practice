package com.github.leetcode.LC2438;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void productQueries() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{2,4,64}, solution.productQueries(15, new int[][]{{0,1}, {2,2}, {0,3}}));
        assertArrayEquals(new int[]{2}, solution.productQueries(2, new int[][]{{0, 0}}));
    }
}
