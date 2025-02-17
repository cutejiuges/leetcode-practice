package com.github.leetcode.LC1287;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestFindSpecialInteger() {
        Solution solution = new Solution();
        assertEquals(6, solution.findSpecialInteger(new int[]{1,2,2,6,6,6,6,7,10}));
    }
}
