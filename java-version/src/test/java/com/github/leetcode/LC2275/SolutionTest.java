package com.github.leetcode.LC2275;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestLargestCombiantion() {
        Solution solution = new Solution();
        assertEquals(4, solution.largestCombination(new int[]{16,17,71,62,12,24,14}));
        assertEquals(2, solution.largestCombination(new int[]{8, 8}));
    }
}
