package com.github.leetcode.LC2270;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestWaysToSplitArray() {
        Solution solution = new Solution();
        assertEquals(2, solution.waysToSplitArray(new int[]{10,4,-8,7}));
        assertEquals(2, solution.waysToSplitArray(new int[]{2,3,1,0}));
    }
}
