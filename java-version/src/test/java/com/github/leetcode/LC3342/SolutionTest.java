package com.github.leetcode.LC3342;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minTimeToReach() {
        Solution solution = new Solution();
        assertEquals(7, solution.minTimeToReach(new int[][]{{0,4}, {4,4}}));
        assertEquals(6, solution.minTimeToReach(new int[][]{{0,0,0,0}, {0,0,0,0}}));
        assertEquals(4, solution.minTimeToReach(new int[][]{{0,1}, {1,2}}));
        assertEquals(71, solution.minTimeToReach(new int[][]{{0,58}, {27,69}}));
    }
}
