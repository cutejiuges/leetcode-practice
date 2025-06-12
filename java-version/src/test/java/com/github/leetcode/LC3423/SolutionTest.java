package com.github.leetcode.LC3423;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxAdjacentDistance() {
        Solution solution = new Solution();
        assertEquals(3, solution.maxAdjacentDistance(new int[]{1,2,4}));
        assertEquals(3, solution.maxAdjacentDistance(new int[]{-3, 0}));
        assertEquals(5, solution.maxAdjacentDistance(new int[]{-5,-10,-5}));
    }
}
