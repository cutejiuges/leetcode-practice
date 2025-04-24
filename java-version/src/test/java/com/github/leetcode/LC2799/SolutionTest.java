package com.github.leetcode.LC2799;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countCompleteSubarrays() {
        Solution solution = new Solution();
        assertEquals(4, solution.countCompleteSubarrays(new int[]{1,3,1,2,2}));
        assertEquals(10, solution.countCompleteSubarrays(new int[]{5,5,5,5}));
    }
}
