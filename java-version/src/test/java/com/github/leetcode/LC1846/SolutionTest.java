package com.github.leetcode.LC1846;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maximumElementAfterDecrementingAndRearranging() {
        Solution solution = new Solution();
        assertEquals(2, solution.maximumElementAfterDecrementingAndRearranging(new int[] {2,2,1,2,1}));
        assertEquals(3, solution.maximumElementAfterDecrementingAndRearranging(new int[] {100,1,1000}));
        assertEquals(5, solution.maximumElementAfterDecrementingAndRearranging(new int[] {1,2,3,4,5}));
    }
}
