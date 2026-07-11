package com.github.leetcode.LC2685;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countCompleteComponents() {
        Solution solution = new Solution();
        assertEquals(3, solution.countCompleteComponents(6, new int[][]{{0,1},{0,2},{1,2},{3,4}}));
        assertEquals(1, solution.countCompleteComponents(6, new int[][]{{0,1},{0,2},{1,2},{3,4},{3,5}}));
    }
}
