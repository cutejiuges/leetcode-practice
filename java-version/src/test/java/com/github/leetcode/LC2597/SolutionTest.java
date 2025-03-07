package com.github.leetcode.LC2597;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void beautifulSubsets() {
        Solution solution = new Solution();
        assertEquals(4, solution.beautifulSubsets(new int[]{2,4,6}, 2));
        assertEquals(1, solution.beautifulSubsets(new int[]{1}, 1));
    }
}
