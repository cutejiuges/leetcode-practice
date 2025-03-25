package com.github.leetcode.LC2711;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void differenceOfDistinctValues() {
        Solution solution = new Solution();
        assertArrayEquals(new int[][]{{1,1,0},{1,0,1},{0,1,1}}, solution.differenceOfDistinctValues(new int[][]{{1,2,3},{3,1,5},{3,2,1}}));
        assertArrayEquals(new int[][]{{0}}, solution.differenceOfDistinctValues(new int[][]{{1}}));
    }
}
