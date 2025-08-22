package com.github.leetcode.LC3195;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minimumArea() {
        Solution solution = new Solution();
        assertEquals(6, solution.minimumArea(new int[][]{{0,1,0}, {1,0,1}}));
        assertEquals(1, solution.minimumArea(new int[][]{{0, 1}, {0, 0}}));
    }
}
