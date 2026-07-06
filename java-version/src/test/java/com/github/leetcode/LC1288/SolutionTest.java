package com.github.leetcode.LC1288;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void removeCoveredIntervals() {
        Solution solution = new Solution();
        assertEquals(2, solution.removeCoveredIntervals(new int[][]{{1,4}, {3,6}, {2,8}}));
        assertEquals(1, solution.removeCoveredIntervals(new int[][]{{1,2}, {1,4}, {3,4}}));
    }
}
