package com.github.leetcode.LC1590;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minSubarray() {
        Solution solution = new Solution();
        assertEquals(1, solution.minSubarray(new int[]{3,1,4,2}, 6));
        assertEquals(2, solution.minSubarray(new int[]{6,3,5,2}, 9));
        assertEquals(0, solution.minSubarray(new int[]{1,2,3}, 3));
        assertEquals(-1, solution.minSubarray(new int[]{1,2,3}, 7));
        assertEquals(0, solution.minSubarray(new int[]{1000000000,1000000000,1000000000}, 3));
    }
}
