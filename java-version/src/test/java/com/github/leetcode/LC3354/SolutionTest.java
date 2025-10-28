package com.github.leetcode.LC3354;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countValidSelections() {
        Solution solution = new Solution();
        assertEquals(2, solution.countValidSelections(new int[]{1,0,2,0,3}));
        assertEquals(0, solution.countValidSelections(new int[]{2,3,4,0,4,1,0}));
    }
}
