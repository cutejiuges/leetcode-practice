package com.github.leetcode.LC2943;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximizeSquareHoleArea() {
        Solution solution = new Solution();
        assertEquals(4, solution.maximizeSquareHoleArea(2, 1, new int[]{2,3}, new int[]{2}));
        assertEquals(4, solution.maximizeSquareHoleArea(1, 1, new int[]{2}, new int[]{2}));
        assertEquals(4, solution.maximizeSquareHoleArea(2, 3, new int[]{2,3}, new int[]{2,4}));
    }
}
