package com.github.leetcode.LC2962;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countSubarrays() {
        Solution solution = new Solution();
        assertEquals(6, solution.countSubarrays(new int[]{1,3,2,3,3}, 2));
        assertEquals(0, solution.countSubarrays(new int[]{1,4,2,1}, 3));
    }
}
