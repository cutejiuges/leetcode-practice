package com.github.leetcode.LC2492;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minScore() {
        Solution solution = new Solution();
        assertEquals(5, solution.minScore(4, new int[][]{{1,2,9}, {2,3,6}, {2,4,5}, {1,4,7}}));
        assertEquals(2, solution.minScore(4, new int[][]{{1,2,2}, {1,3,4}, {3,4,7}}));
    }
}
