package com.github.leetcode.LC1432;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxDiff() {
        Solution solution = new Solution();
        assertEquals(888, solution.maxDiff(555));
        assertEquals(8, solution.maxDiff(9));
        assertEquals(820000, solution.maxDiff(123456));
        assertEquals(80000, solution.maxDiff(10000));
        assertEquals(8700, solution.maxDiff(9288));
    }
}
