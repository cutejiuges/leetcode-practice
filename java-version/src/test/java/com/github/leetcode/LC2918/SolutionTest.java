package com.github.leetcode.LC2918;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minSum() {
        Solution solution = new Solution();
        assertEquals(12, solution.minSum(new int[]{3,2,0,1,0}, new int[]{6,5,0}));
        assertEquals(-1, solution.minSum(new int[]{2,0,2,0}, new int[]{1,4}));
    }
}
