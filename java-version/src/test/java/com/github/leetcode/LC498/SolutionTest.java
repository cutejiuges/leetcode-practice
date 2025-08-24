package com.github.leetcode.LC498;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findDiagonalOrder() {
        Solution solution = new Solution();
        assertArrayEquals(new int[]{1,2,4,7,5,3,6,8,9}, solution.findDiagonalOrder(new int[][]{{1,2,3}, {4,5,6}, {7,8,9}}));
        assertArrayEquals(new int[]{1,2,3,4}, solution.findDiagonalOrder(new int[][]{{1,2}, {3,4}}));
    }
}
