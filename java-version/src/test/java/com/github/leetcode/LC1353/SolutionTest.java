package com.github.leetcode.LC1353;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxEvents() {
        Solution solution = new Solution();
        assertEquals(3, solution.maxEvents(new int[][]{{1,2}, {2,3}, {3,4}}));
        assertEquals(4, solution.maxEvents(new int[][]{{1,2}, {2,3}, {3,4}, {1,2}}));
    }
}
