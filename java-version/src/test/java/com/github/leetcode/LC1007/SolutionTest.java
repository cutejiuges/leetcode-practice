package com.github.leetcode.LC1007;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minDominoRotations() {
        Solution solution = new Solution();
        assertEquals(2, solution.minDominoRotations(new int[]{2,1,2,4,2,2}, new int[]{5,2,6,2,3,2}));
        assertEquals(-1, solution.minDominoRotations(new int[]{3,5,1,2,3}, new int[]{3,6,3,3,4}));
    }
}
