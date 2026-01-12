package com.github.leetcode.LC1266;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minTimeToVisitAllPoints() {
        Solution solution = new Solution();
        assertEquals(7, solution.minTimeToVisitAllPoints(new int[][]{{1,1}, {3,4}, {-1,0}}));
        assertEquals(5, solution.minTimeToVisitAllPoints(new int[][]{{3,2}, {-2,2}}));
    }
}
