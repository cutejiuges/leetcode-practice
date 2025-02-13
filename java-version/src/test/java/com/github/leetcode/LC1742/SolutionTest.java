package com.github.leetcode.LC1742;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestCountBall() {
        Solution solution = new Solution();
        assertEquals(2, solution.countBalls(1, 10));
        assertEquals(2, solution.countBalls(5, 15));
        assertEquals(2, solution.countBalls(19, 28));
    }
}
