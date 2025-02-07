package com.github.leetcode.LC3065;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMinOperations() {
        Solution solution = new Solution();
        assertEquals(3, solution.minOperations(new int[]{2,11,10,1,3}, 10));
        assertEquals(0, solution.minOperations(new int[]{1,1,2,4,9}, 1));
        assertEquals(4, solution.minOperations(new int[]{1,1,2,4,9}, 9));
    }
}
