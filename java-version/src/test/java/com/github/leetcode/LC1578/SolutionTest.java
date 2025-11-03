package com.github.leetcode.LC1578;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minCost() {
        Solution solution = new Solution();
        assertEquals(3, solution.minCost("abaac", new int[]{1,2,3,4,5}));
        assertEquals(0, solution.minCost("abc", new int[]{1,2,3}));
        assertEquals(2, solution.minCost("aabaa", new int[]{1,2,3,4,1}));
    }
}
