package com.github.leetcode.LC3000;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void areaOfMaxDiagonal() {
        Solution solution = new Solution();
        assertEquals(48, solution.areaOfMaxDiagonal(new int[][]{{9,3}, {6,8}}));
        assertEquals(12, solution.areaOfMaxDiagonal(new int[][]{{3,4}, {4,3}}));
    }
}
