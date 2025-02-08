package com.github.leetcode.LC63;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestUniquePathsWithObstacles() {
        Solution solution = new Solution();
        assertEquals(2, solution.uniquePathsWithObstacles(new int[][]{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}));
        assertEquals(1, solution.uniquePathsWithObstacles(new int[][]{{0, 1}, {0, 0}}));
    }
}
