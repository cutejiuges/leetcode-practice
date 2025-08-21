package com.github.leetcode.LC1504;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numSubmat() {
        Solution solution = new Solution();
        assertEquals(13, solution.numSubmat(new int[][]{{1,0,1}, {1,1,0}, {1,1,0}}));
        assertEquals(24, solution.numSubmat(new int[][]{{0,1,1,0}, {0,1,1,1}, {1,1,1,0}}));
    }
}
