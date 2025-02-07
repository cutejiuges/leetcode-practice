package com.github.leetcode.LC3066;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestMinOperations() {
        Solution solution = new Solution();
        assertEquals(2, solution.minOperations(new int[]{2,11,10,1,3}, 10));
        assertEquals(4, solution.minOperations(new int[]{1,1,2,4,9}, 20));
    }
}
