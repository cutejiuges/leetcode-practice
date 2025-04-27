package com.github.leetcode.LC3392;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countSubarrays() {
        Solution solution = new Solution();
        assertEquals(1, solution.countSubarrays(new int[]{1,2,1,4,1}));
        assertEquals(0, solution.countSubarrays(new int[]{1,1,1}));
    }
}
