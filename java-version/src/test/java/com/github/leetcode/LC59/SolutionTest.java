package com.github.leetcode.LC59;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestGenerateMatrix() {
        Solution solution = new Solution();
        assertArrayEquals(new int[][]{{1,2,3},{8,9,4},{7,6,5}}, solution.generateMatrix(3));
        assertArrayEquals(new int[][]{{1}}, solution.generateMatrix(1));
    }
}
