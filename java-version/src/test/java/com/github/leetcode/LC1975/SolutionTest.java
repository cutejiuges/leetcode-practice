package com.github.leetcode.LC1975;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxMatrixSum() {
        Solution solution = new Solution();
        assertEquals(4, solution.maxMatrixSum(new int[][]{{1,-1}, {-1,1}}));
        assertEquals(16, solution.maxMatrixSum(new int[][]{{1,2,3}, {-1,-2,-3}, {1,2,3}}));
    }
}
