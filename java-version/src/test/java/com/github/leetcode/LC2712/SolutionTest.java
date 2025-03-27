package com.github.leetcode.LC2712;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumCost() {
        Solution solution = new Solution();
        assertEquals(2, solution.minimumCost("0011"));
        assertEquals(9, solution.minimumCost("010101"));
    }
}
