package com.github.leetcode.LC2239;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestFinsClosestNumber() {
        Solution solution = new Solution();
        assertEquals(1, solution.findClosestNumber(new int[]{-4,-2,1,4,8}));
        assertEquals(1, solution.findClosestNumber(new int[]{2,-1,1}));
    }
}
