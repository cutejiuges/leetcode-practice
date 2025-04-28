package com.github.leetcode.LC2302;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countSubarrays() {
        Solution solution = new Solution();
        assertEquals(6, solution.countSubarrays(new int[]{2,1,4,3,5}, 10));
        assertEquals(5, solution.countSubarrays(new int[]{1,1,1}, 5));
        assertEquals(13, solution.countSubarrays(new int[]{5,2,6,8,9,7}, 50));
    }
}
