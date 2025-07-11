package com.github.leetcode.LC3169;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countDays() {
        Solution solution = new Solution();
        assertEquals(2, solution.countDays(10, new int[][]{{5,7},{1,3},{9,10}}));
        assertEquals(1, solution.countDays(5, new int[][]{{2,4},{1,3}}));
        assertEquals(0, solution.countDays(6, new int[][]{{1,6}}));
    }
}
